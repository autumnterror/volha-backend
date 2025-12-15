package api

import (
	"context"
	"github.com/autumnterror/volha-backend/pkg/views"

	"github.com/autumnterror/breezynotes/pkg/log"
	productsRPC "github.com/autumnterror/volha-backend/api/proto/gen"
	"github.com/autumnterror/volha-backend/internal/product-service/domain"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ServerAPI) CreateMaterial(ctx context.Context, req *productsRPC.Material) (*emptypb.Empty, error) {
	const op = "grpc.CreateMaterial"
	log.Blue(op)

	_, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return nil, s.s.Create(ctx, domain.MaterialFromRpc(req), views.Material)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) UpdateMaterial(ctx context.Context, req *productsRPC.Material) (*emptypb.Empty, error) {
	const op = "grpc.UpdateMaterial"
	log.Blue(op)

	_, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return nil, s.s.Update(ctx, domain.MaterialFromRpc(req), views.Material)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) DeleteMaterial(ctx context.Context, req *productsRPC.Id) (*emptypb.Empty, error) {
	const op = "grpc.DeleteMaterial"
	log.Blue(op)

	_, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return nil, s.s.Delete(ctx, req.GetId(), views.Material)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) GetAllMaterials(ctx context.Context, _ *emptypb.Empty) (*productsRPC.MaterialList, error) {
	const op = "grpc.GetAllMaterials"
	log.Blue(op)

	res, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return s.s.GetAll(ctx, views.Material)
	})
	if err != nil {
		return nil, err
	}

	materials := res.([]*domain.Material)
	return &productsRPC.MaterialList{
		Items: domain.MaterialsToRpc(materials),
	}, nil
}

func (s *ServerAPI) GetMaterial(ctx context.Context, req *productsRPC.Id) (*productsRPC.Material, error) {
	const op = "grpc.GetMaterial"
	log.Blue(op)

	res, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return s.s.Get(ctx, req.GetId(), views.Material)
	})
	if err != nil {
		return nil, err
	}

	return domain.MaterialToRpc(res.(*domain.Material)), nil
}
