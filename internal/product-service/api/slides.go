package api

import (
	"context"
	"github.com/autumnterror/volha-backend/pkg/views"

	"github.com/autumnterror/breezynotes/pkg/log"
	productsRPC "github.com/autumnterror/volha-backend/api/proto/gen"
	"github.com/autumnterror/volha-backend/internal/product-service/domain"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ServerAPI) CreateSlide(ctx context.Context, req *productsRPC.Slide) (*emptypb.Empty, error) {
	const op = "grpc.CreateSlide"
	log.Blue(op)

	_, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return nil, s.s.Create(ctx, domain.SlideFromRpc(req), views.Slide)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) UpdateSlide(ctx context.Context, req *productsRPC.Slide) (*emptypb.Empty, error) {
	const op = "grpc.UpdateSlide"
	log.Blue(op)

	_, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return nil, s.s.Update(ctx, domain.SlideFromRpc(req), views.Slide)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) DeleteSlide(ctx context.Context, req *productsRPC.Id) (*emptypb.Empty, error) {
	const op = "grpc.DeleteSlide"
	log.Blue(op)

	_, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return nil, s.s.Delete(ctx, req.GetId(), views.Slide)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) GetAllSlides(ctx context.Context, _ *emptypb.Empty) (*productsRPC.SlideList, error) {
	const op = "grpc.GetAllSlides"
	log.Blue(op)

	res, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return s.s.GetAll(ctx, views.Slide)
	})
	if err != nil {
		return nil, err
	}

	slides := res.([]*domain.Slide)
	return &productsRPC.SlideList{
		Items: domain.SlidesToRpc(slides),
	}, nil
}

func (s *ServerAPI) GetSlide(ctx context.Context, req *productsRPC.Id) (*productsRPC.Slide, error) {
	const op = "grpc.GetSlide"
	log.Blue(op)

	res, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return s.s.Get(ctx, req.GetId(), views.Slide)
	})
	if err != nil {
		return nil, err
	}

	return domain.SlideToRpc(res.(*domain.Slide)), nil
}
