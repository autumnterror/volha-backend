package service

import (
	"context"
	"github.com/autumnterror/volha-backend/pkg/views"
)

func (s *ProductsService) GetAll(ctx context.Context, _type views.Type) (any, error) {
	const op = "service.GetAll"
	if err := validateBasicType(_type); err != nil {
		return nil, wrapServiceCheck(op, err)
	}
	repo, err := s.basicRepo(ctx)
	if err != nil {
		return nil, wrapServiceCheck(op, err)
	}

	return repo.GetAll(ctx, _type)
}

func (s *ProductsService) Get(ctx context.Context, id string, _type views.Type) (any, error) {
	const op = "service.Get"
	if err := validateBasicType(_type); err != nil {
		return nil, wrapServiceCheck(op, err)
	}
	if err := validateID(id); err != nil {
		return nil, wrapServiceCheck(op, err)
	}
	repo, err := s.basicRepo(ctx)
	if err != nil {
		return nil, wrapServiceCheck(op, err)
	}

	return repo.Get(ctx, id, _type)
}

func (s *ProductsService) Create(ctx context.Context, obj any, _type views.Type) error {
	const op = "service.Create"
	if err := validateBasicType(_type); err != nil {
		return wrapServiceCheck(op, err)
	}
	if err := validateBasicPayload(obj, _type); err != nil {
		return wrapServiceCheck(op, err)
	}
	return s.runInTx(ctx, op, func(ctx context.Context) error {
		repo, err := s.basicRepo(ctx)
		if err != nil {
			return wrapServiceCheck(op, err)
		}
		return repo.Create(ctx, obj, _type)
	})
}

func (s *ProductsService) Update(ctx context.Context, obj any, _type views.Type) error {
	const op = "service.Update"
	if err := validateBasicType(_type); err != nil {
		return wrapServiceCheck(op, err)
	}
	if err := validateBasicPayload(obj, _type); err != nil {
		return wrapServiceCheck(op, err)
	}
	return s.runInTx(ctx, op, func(ctx context.Context) error {
		repo, err := s.basicRepo(ctx)
		if err != nil {
			return wrapServiceCheck(op, err)
		}
		return repo.Update(ctx, obj, _type)
	})
}

func (s *ProductsService) Delete(ctx context.Context, id string, _type views.Type) error {
	const op = "service.Delete"
	if err := validateBasicType(_type); err != nil {
		return wrapServiceCheck(op, err)
	}
	if err := validateID(id); err != nil {
		return wrapServiceCheck(op, err)
	}
	return s.runInTx(ctx, op, func(ctx context.Context) error {
		repo, err := s.basicRepo(ctx)
		if err != nil {
			return wrapServiceCheck(op, err)
		}
		return repo.Delete(ctx, id, _type)
	})
}
