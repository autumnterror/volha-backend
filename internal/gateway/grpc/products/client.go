package products

import (
	"context"
	"fmt"
	"github.com/autumnterror/breezynotes/pkg/utils/format"
	"github.com/autumnterror/volha-backend/api/proto/gen"
	"github.com/autumnterror/volha-backend/internal/gateway/config"

	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

type Client struct {
	API productsRPC.ProductsClient
}

func New(
	cfg *config.Config,
) (*Client, error) {
	const op = "grpc.products.New"

	retryOpts := []grpcretry.CallOption{
		grpcretry.WithCodes(codes.Aborted, codes.DeadlineExceeded),
		grpcretry.WithMax(uint(cfg.RetriesCount)),
		grpcretry.WithPerRetryTimeout(cfg.Timeout),
		grpcretry.WithBackoff(grpcretry.BackoffExponential(cfg.Backoff)),
	}

	cc, err := grpc.NewClient(
		cfg.AddrProducts,
		grpc.WithChainUnaryInterceptor(
			grpcretry.UnaryClientInterceptor(retryOpts...),
		),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, format.Error(op, err)
	}

	log.Println("start connection...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cc.Connect()
	for {
		state := cc.GetState()
		if state == connectivity.Ready {
			log.Println("CONNECT!!")
			break
		}

		if !cc.WaitForStateChange(ctx, state) {
			return nil, fmt.Errorf("connection timeout")
		}
	}

	return &Client{
		API: productsRPC.NewProductsClient(cc),
	}, nil
}
