package api

import (
	"context"
	"github.com/autumnterror/volha-backend/pkg/views"

	"github.com/autumnterror/breezynotes/pkg/log"
	productsRPC "github.com/autumnterror/volha-backend/api/proto/gen"
	"github.com/autumnterror/volha-backend/internal/product-service/domain"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ServerAPI) CreateBrand(ctx context.Context, req *productsRPC.Brand) (*emptypb.Empty, error) {
	const op = "grpc.CreateBrand"
	log.Blue(op)

	_, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return nil, s.s.Create(ctx, domain.BrandFromRpc(req), views.Brand)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) UpdateBrand(ctx context.Context, req *productsRPC.Brand) (*emptypb.Empty, error) {
	const op = "grpc.UpdateBrand"
	log.Blue(op)

	_, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return nil, s.s.Update(ctx, domain.BrandFromRpc(req), views.Brand)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) DeleteBrand(ctx context.Context, req *productsRPC.Id) (*emptypb.Empty, error) {
	const op = "grpc.DeleteBrand"
	log.Blue(op)

	_, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return nil, s.s.Delete(ctx, req.GetId(), views.Brand)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) GetAllBrands(ctx context.Context, _ *emptypb.Empty) (*productsRPC.BrandList, error) {
	const op = "grpc.GetAllBrands"
	log.Blue(op)

	res, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return s.s.GetAll(ctx, views.Brand)
	})
	if err != nil {
		return nil, err
	}

	brands := res.([]*domain.Brand)
	return &productsRPC.BrandList{
		Items: domain.BrandsToRpc(brands),
	}, nil
}

func (s *ServerAPI) GetBrand(ctx context.Context, req *productsRPC.Id) (*productsRPC.Brand, error) {
	const op = "grpc.GetBrand"
	log.Blue(op)

	res, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return s.s.Get(ctx, req.GetId(), views.Brand)
	})
	if err != nil {
		return nil, err
	}

	return domain.BrandToRpc(res.(*domain.Brand)), nil
}
