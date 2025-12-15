package api

import (
	"context"
	"github.com/autumnterror/volha-backend/pkg/views"

	"github.com/autumnterror/breezynotes/pkg/log"
	productsRPC "github.com/autumnterror/volha-backend/api/proto/gen"
	"github.com/autumnterror/volha-backend/internal/product-service/domain"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ServerAPI) CreateColor(ctx context.Context, req *productsRPC.Color) (*emptypb.Empty, error) {
	const op = "grpc.CreateColor"
	log.Blue(op)

	_, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return nil, s.s.Create(ctx, domain.ColorFromRpc(req), views.Color)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) UpdateColor(ctx context.Context, req *productsRPC.Color) (*emptypb.Empty, error) {
	const op = "grpc.UpdateColor"
	log.Blue(op)

	_, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return nil, s.s.Update(ctx, domain.ColorFromRpc(req), views.Color)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) DeleteColor(ctx context.Context, req *productsRPC.Id) (*emptypb.Empty, error) {
	const op = "grpc.DeleteColor"
	log.Blue(op)

	_, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return nil, s.s.Delete(ctx, req.GetId(), views.Color)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) GetAllColors(ctx context.Context, _ *emptypb.Empty) (*productsRPC.ColorList, error) {
	const op = "grpc.GetAllColors"
	log.Blue(op)

	res, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return s.s.GetAll(ctx, views.Color)
	})
	if err != nil {
		return nil, err
	}

	colors := res.([]*domain.Color)
	return &productsRPC.ColorList{
		Items: domain.ColorsToRpc(colors),
	}, nil
}

func (s *ServerAPI) GetColor(ctx context.Context, req *productsRPC.Id) (*productsRPC.Color, error) {
	const op = "grpc.GetColor"
	log.Blue(op)

	res, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return s.s.Get(ctx, req.GetId(), views.Color)
	})
	if err != nil {
		return nil, err
	}

	return domain.ColorToRpc(res.(*domain.Color)), nil
}
