package grpc

import (
	"context"
	productsRPC "github.com/autumnterror/volha-proto/gen/products"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"productService/internal/pkg/psql"
	"productService/internal/utils/format"
	"productService/internal/views"
)

type ServerAPI struct {
	productsRPC.UnimplementedProductsServer
	API psql.ProductRepo
}

func Register(
	server *grpc.Server,
	API psql.ProductRepo,
) {
	productsRPC.RegisterProductsServer(server, &ServerAPI{
		API: API,
	})
}

func convertToProductView(p *productsRPC.Product) *views.Product {
	return &views.Product{
		Id:          p.Id,
		Title:       p.Title,
		Article:     p.Article,
		Brand:       p.Brand,
		Country:     p.Country,
		Width:       int(p.Width),
		Height:      int(p.Height),
		Depth:       int(p.Depth),
		Materials:   p.Materials,
		Colors:      p.Color,
		Photos:      p.Photos,
		Seems:       p.Seems,
		Price:       int(p.Price),
		Description: p.Description,
	}
}

func (s *ServerAPI) Create(
	ctx context.Context,
	req *productsRPC.Product,
) (*emptypb.Empty, error) {
	const op = "productsRPC.ServerAPI.Create"

	type result struct {
		err error
	}
	res := make(chan result, 1)

	go func() {
		p := convertToProductView(req)

		if err := s.API.Create(p); err != nil {
			res <- result{err: format.Error(op, status.Error(codes.Internal, err.Error()))}
			log.Println(format.Error(op, status.Error(codes.Internal, err.Error())))
			return
		}
		log.Printf("Create %s\n", format.Struct(p))
		res <- result{err: nil}
	}()

	select {
	case <-ctx.Done():
		log.Println(format.Error(op, status.Error(codes.DeadlineExceeded, "Context dead")))
		return nil, format.Error(op, status.Error(codes.DeadlineExceeded, "Context dead"))
	case r := <-res:
		return nil, r.err
	}
}

func (s *ServerAPI) Update(
	ctx context.Context,
	req *productsRPC.Product,
) (*emptypb.Empty, error) {
	const op = "productsRPC.ServerAPI.Update"

	type result struct {
		err error
	}
	res := make(chan result, 1)

	go func() {
		p := convertToProductView(req)

		if err := s.API.Update(p, p.Id); err != nil {
			res <- result{err: format.Error(op, status.Error(codes.Internal, err.Error()))}
			log.Println(format.Error(op, status.Error(codes.Internal, err.Error())))
			return
		}
		log.Printf("Update %s: %s\n", p.Id, format.Struct(p))
		res <- result{err: nil}
	}()

	select {
	case <-ctx.Done():
		log.Println(format.Error(op, status.Error(codes.DeadlineExceeded, "Context dead")))
		return nil, format.Error(op, status.Error(codes.DeadlineExceeded, "Context dead"))
	case r := <-res:
		return nil, r.err
	}
}

func (s *ServerAPI) Delete(
	ctx context.Context,
	req *productsRPC.Id,
) (*emptypb.Empty, error) {
	const op = "productsRPC.ServerAPI.Delete"

	type result struct {
		err error
	}
	res := make(chan result, 1)

	go func() {
		if err := s.API.Delete(req.Id); err != nil {
			res <- result{err: format.Error(op, status.Error(codes.Internal, err.Error()))}
			log.Println(format.Error(op, status.Error(codes.Internal, err.Error())))
			return
		}
		log.Printf("Delete %s\n", req.Id)
		res <- result{err: nil}
	}()

	select {
	case <-ctx.Done():
		log.Println(format.Error(op, status.Error(codes.DeadlineExceeded, "Context dead")))
		return nil, format.Error(op, status.Error(codes.DeadlineExceeded, "Context dead"))
	case r := <-res:
		return nil, r.err
	}
}
