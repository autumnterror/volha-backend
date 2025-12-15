package api

import (
	"context"

	"github.com/autumnterror/breezynotes/pkg/log"
	productsRPC "github.com/autumnterror/volha-backend/api/proto/gen"
	"github.com/autumnterror/volha-backend/internal/product-service/domain"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ServerAPI) CreateProductColorPhotos(ctx context.Context, req *productsRPC.ProductColorPhotos) (*emptypb.Empty, error) {
	const op = "grpc.CreateProductColorPhotos"
	log.Blue(op)

	_, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return nil, s.s.CreateProductColorPhotos(ctx, domain.ProductColorPhotosFromRpc(req))
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) UpdateProductColorPhotos(ctx context.Context, req *productsRPC.ProductColorPhotos) (*emptypb.Empty, error) {
	const op = "grpc.UpdateProductColorPhotos"
	log.Blue(op)

	_, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return nil, s.s.UpdateProductColorPhotos(ctx, domain.ProductColorPhotosFromRpc(req))
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) DeleteProductColorPhotos(ctx context.Context, req *productsRPC.ProductColorPhotosId) (*emptypb.Empty, error) {
	const op = "grpc.DeleteProductColorPhotos"
	log.Blue(op)

	_, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return nil, s.s.DeleteProductColorPhotos(ctx, req.GetProductId(), req.GetColorId())
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) GetAllProductColorPhotos(ctx context.Context, _ *emptypb.Empty) (*productsRPC.ProductColorPhotosList, error) {
	const op = "grpc.GetAllProductColorPhotos"
	log.Blue(op)

	res, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return s.s.GetAllProductColorPhotos(ctx)
	})
	if err != nil {
		return nil, err
	}

	pcps := res.([]domain.ProductColorPhotos)
	return &productsRPC.ProductColorPhotosList{
		Items: domain.ProductColorPhotosListToRpc(pcps),
	}, nil
}

func (s *ServerAPI) GetPhotosByProductAndColor(ctx context.Context, req *productsRPC.ProductColorPhotosId) (*productsRPC.PhotoList, error) {
	const op = "grpc.GetPhotosByProductAndColor"
	log.Blue(op)

	res, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return s.s.GetPhotosByProductColor(ctx, req.GetProductId(), req.GetColorId())
	})
	if err != nil {
		return nil, err
	}

	photos := res.([]string)
	return &productsRPC.PhotoList{
		Items: photos,
	}, nil
}
