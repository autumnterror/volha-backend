package api

import (
	"context"
	"github.com/autumnterror/volha-backend/pkg/views"

	"github.com/autumnterror/breezynotes/pkg/log"
	productsRPC "github.com/autumnterror/volha-backend/api/proto/gen"
	"github.com/autumnterror/volha-backend/internal/product-service/domain"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ServerAPI) CreateCategory(ctx context.Context, req *productsRPC.Category) (*emptypb.Empty, error) {
	const op = "grpc.CreateCategory"
	log.Blue(op)

	_, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return nil, s.s.Create(ctx, domain.CategoryFromRpc(req), views.Category)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) UpdateCategory(ctx context.Context, req *productsRPC.Category) (*emptypb.Empty, error) {
	const op = "grpc.UpdateCategory"
	log.Blue(op)

	_, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return nil, s.s.Update(ctx, domain.CategoryFromRpc(req), views.Category)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) DeleteCategory(ctx context.Context, req *productsRPC.Id) (*emptypb.Empty, error) {
	const op = "grpc.DeleteCategory"
	log.Blue(op)

	_, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return nil, s.s.Delete(ctx, req.GetId(), views.Category)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) GetAllCategories(ctx context.Context, _ *emptypb.Empty) (*productsRPC.CategoryList, error) {
	const op = "grpc.GetAllCategories"
	log.Blue(op)

	res, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return s.s.GetAll(ctx, views.Category)
	})
	if err != nil {
		return nil, err
	}

	categories := res.([]*domain.Category)
	return &productsRPC.CategoryList{
		Items: domain.CategoriesToRpc(categories),
	}, nil
}

func (s *ServerAPI) GetCategory(ctx context.Context, req *productsRPC.Id) (*productsRPC.Category, error) {
	const op = "grpc.GetCategory"
	log.Blue(op)

	res, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return s.s.Get(ctx, req.GetId(), views.Category)
	})
	if err != nil {
		return nil, err
	}

	return domain.CategoryToRpc(res.(*domain.Category)), nil
}
