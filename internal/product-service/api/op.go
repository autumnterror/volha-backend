package api

import (
	"context"
	"errors"
	"github.com/autumnterror/breezynotes/pkg/log"
	"github.com/autumnterror/volha-backend/api/proto/gen"
	"github.com/autumnterror/volha-backend/internal/product-service/domain"
	"github.com/autumnterror/volha-backend/internal/product-service/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ServerAPI struct {
	productsRPC.UnimplementedProductsServer
	s *service.ProductsService
}

func Register(
	server *grpc.Server,
	s *service.ProductsService,
) {
	productsRPC.RegisterProductsServer(server, &ServerAPI{
		s: s,
	})
}

func handleCRUDResponse(ctx context.Context, op string, action func() (any, error)) (any, error) {
	type resV struct {
		res any
		err error
	}
	res := make(chan resV, 1)
	go func() {
		r, err := action()
		res <- resV{r, err}
	}()
	select {
	case <-ctx.Done():
		log.Error(op, "Context dead", ctx.Err())
		return nil, status.Error(codes.DeadlineExceeded, "Context dead")
	case r := <-res:
		if r.err != nil {
			log.Error(op, "", r.err)
			switch {
			case errors.Is(r.err, domain.ErrUnknownType):
				return nil, status.Error(codes.Unimplemented, r.err.Error())
			case errors.Is(r.err, domain.ErrInvalidType):
				return nil, status.Error(codes.InvalidArgument, r.err.Error())
			case errors.Is(r.err, domain.ErrNotFound):
				return nil, status.Error(codes.NotFound, r.err.Error())
			case errors.Is(r.err, domain.ErrAlreadyExists):
				return nil, status.Error(codes.AlreadyExists, r.err.Error())
			case errors.Is(r.err, domain.ErrForeignKey):
				return nil, status.Error(codes.FailedPrecondition, r.err.Error())
			case errors.Is(r.err, service.ErrBadServiceCheck):
				return nil, status.Error(codes.InvalidArgument, r.err.Error())
			default:
				return nil, status.Error(codes.Internal, "check logs")
			}
		}
		log.Green(op)
		if r.res != nil {
			return r.res, nil
		}
		return nil, nil
	}
}
