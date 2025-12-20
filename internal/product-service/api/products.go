package api

import (
	"context"

	"github.com/autumnterror/breezynotes/pkg/log"
	productsRPC "github.com/autumnterror/volha-backend/api/proto/gen"
	"github.com/autumnterror/volha-backend/internal/product-service/domain"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ServerAPI) SearchProducts(ctx context.Context, req *productsRPC.ProductSearchWithPagination) (*productsRPC.ProductList, error) {
	const op = "grpc.SearchProducts"
	log.Blue(op)

	res, err := handleCRUDResponse(ctx, op, func() (any, error) {
		pr, t, err := s.s.SearchProducts(ctx, domain.ProductSearchFromRpc(req))
		return &productsRPC.ProductList{
			Items: domain.ProductsToRpc(pr),
			Total: int32(t),
		}, err
	})
	if err != nil {
		return nil, err
	}

	return res.(*productsRPC.ProductList), nil
}

func (s *ServerAPI) FilterProducts(ctx context.Context, req *productsRPC.ProductFilter) (*productsRPC.ProductList, error) {
	const op = "grpc.FilterProducts"
	log.Blue(op)

	res, err := handleCRUDResponse(ctx, op, func() (any, error) {
		pr, t, err := s.s.FilterProducts(ctx, domain.ProductFilterFromRpc(req))
		return &productsRPC.ProductList{
			Items: domain.ProductsToRpc(pr),
			Total: int32(t),
		}, err
	})
	if err != nil {
		return nil, err
	}

	return res.(*productsRPC.ProductList), nil
}

func (s *ServerAPI) GetAllProducts(ctx context.Context, req *productsRPC.Pagination) (*productsRPC.ProductList, error) {
	const op = "grpc.GetAllProducts"
	log.Blue(op)

	res, err := handleCRUDResponse(ctx, op, func() (any, error) {
		pr, t, err := s.s.GetAllProducts(ctx, int(req.GetStart()), int(req.GetFinish()))
		return &productsRPC.ProductList{
			Items: domain.ProductsToRpc(pr),
			Total: int32(t),
		}, err
	})
	if err != nil {
		return nil, err
	}

	return res.(*productsRPC.ProductList), nil
}

func (s *ServerAPI) GetProduct(ctx context.Context, req *productsRPC.Id) (*productsRPC.Product, error) {
	const op = "grpc.GetProduct"
	log.Blue(op)

	res, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return s.s.GetProduct(ctx, req.GetId())
	})
	if err != nil {
		return nil, err
	}

	return domain.ProductToRpc(res.(*domain.Product)), nil
}

func (s *ServerAPI) CreateProduct(ctx context.Context, req *productsRPC.ProductId) (*emptypb.Empty, error) {
	const op = "grpc.CreateProduct"
	log.Blue(op)

	_, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return nil, s.s.CreateProduct(ctx, domain.ProductIdFromRpc(req))
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) UpdateProduct(ctx context.Context, req *productsRPC.ProductId) (*emptypb.Empty, error) {
	const op = "grpc.UpdateProduct"
	log.Blue(op)

	_, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return nil, s.s.UpdateProduct(ctx, domain.ProductIdFromRpc(req))
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) DeleteProduct(ctx context.Context, req *productsRPC.Id) (*emptypb.Empty, error) {
	const op = "grpc.DeleteProduct"
	log.Blue(op)

	_, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return nil, s.s.DeleteProduct(ctx, req.GetId())
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}
