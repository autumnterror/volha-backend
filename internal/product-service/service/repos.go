package service

import (
	"context"
	"errors"

	"github.com/autumnterror/volha-backend/internal/product-service/repository"
)

func (s *ProductsService) runInTx(ctx context.Context, op string, fn func(ctx context.Context) error) error {
	if s.tx == nil {
		return wrapServiceCheck(op, errors.New("tx runner is nil"))
	}
	return s.tx.RunInTx(ctx, fn)
}

func (s *ProductsService) productRepo(ctx context.Context) (repository.ProductRepo, error) {
	repo, err := s.repoGetter(ctx, func(p repository.Provider) any {
		return p.Product(ctx)
	})
	if err != nil {
		return nil, err
	}
	res, _ := repo.(repository.ProductRepo)
	if res == nil {
		return nil, errors.New("product repository is nil")
	}
	return res, nil
}

func (s *ProductsService) basicRepo(ctx context.Context) (repository.BasicRepo, error) {
	repo, err := s.repoGetter(ctx, func(p repository.Provider) any {
		return p.Basic(ctx)
	})
	if err != nil {
		return nil, err
	}
	res, _ := repo.(repository.BasicRepo)
	if res == nil {
		return nil, errors.New("basic repository is nil")
	}
	return res, nil
}

func (s *ProductsService) dictRepo(ctx context.Context) (repository.DictRepo, error) {
	repo, err := s.repoGetter(ctx, func(p repository.Provider) any {
		return p.Dict(ctx)
	})
	if err != nil {
		return nil, err
	}
	res, _ := repo.(repository.DictRepo)
	if res == nil {
		return nil, errors.New("dictionary repository is nil")
	}
	return res, nil
}

func (s *ProductsService) pcpRepo(ctx context.Context) (repository.ProductColorPhotosRepo, error) {
	repo, err := s.repoGetter(ctx, func(p repository.Provider) any {
		return p.Pcp(ctx)
	})
	if err != nil {
		return nil, err
	}
	res, _ := repo.(repository.ProductColorPhotosRepo)
	if res == nil {
		return nil, errors.New("product color photos repository is nil")
	}
	return res, nil
}

func (s *ProductsService) repoGetter(ctx context.Context, getter func(repository.Provider) any) (any, error) {
	if s.repos == nil {
		return nil, errors.New("repository provider is nil")
	}
	if getter == nil {
		return nil, errors.New("repository provider is nil")
	}
	return getter(s.repos), nil
}
