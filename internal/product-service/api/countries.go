package api

import (
	"context"
	"github.com/autumnterror/volha-backend/pkg/views"

	"github.com/autumnterror/breezynotes/pkg/log"
	productsRPC "github.com/autumnterror/volha-backend/api/proto/gen"
	"github.com/autumnterror/volha-backend/internal/product-service/domain"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ServerAPI) CreateCountry(ctx context.Context, req *productsRPC.Country) (*emptypb.Empty, error) {
	const op = "grpc.CreateCountry"
	log.Blue(op)

	_, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return nil, s.s.Create(ctx, domain.CountryFromRpc(req), views.Country)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) UpdateCountry(ctx context.Context, req *productsRPC.Country) (*emptypb.Empty, error) {
	const op = "grpc.UpdateCountry"
	log.Blue(op)

	_, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return nil, s.s.Update(ctx, domain.CountryFromRpc(req), views.Country)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) DeleteCountry(ctx context.Context, req *productsRPC.Id) (*emptypb.Empty, error) {
	const op = "grpc.DeleteCountry"
	log.Blue(op)

	_, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return nil, s.s.Delete(ctx, req.GetId(), views.Country)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) GetAllCountries(ctx context.Context, _ *emptypb.Empty) (*productsRPC.CountryList, error) {
	const op = "grpc.GetAllCountries"
	log.Blue(op)

	res, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return s.s.GetAll(ctx, views.Country)
	})
	if err != nil {
		return nil, err
	}

	countries := res.([]*domain.Country)
	return &productsRPC.CountryList{
		Items: domain.CountriesToRpc(countries),
	}, nil
}

func (s *ServerAPI) GetCountry(ctx context.Context, req *productsRPC.Id) (*productsRPC.Country, error) {
	const op = "grpc.GetCountry"
	log.Blue(op)

	res, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return s.s.Get(ctx, req.GetId(), views.Country)
	})
	if err != nil {
		return nil, err
	}

	return domain.CountryToRpc(res.(*domain.Country)), nil
}
