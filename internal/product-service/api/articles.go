package api

import (
	"context"
	"github.com/autumnterror/breezynotes/pkg/log"
	productsRPC "github.com/autumnterror/volha-backend/api/proto/gen"
	"github.com/autumnterror/volha-backend/internal/product-service/domain"
	"github.com/autumnterror/volha-backend/pkg/views"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ServerAPI) CreateArticle(ctx context.Context, req *productsRPC.Article) (*emptypb.Empty, error) {
	const op = "grpc.CreateArticle"
	log.Blue(op)

	_, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return nil, s.s.Create(ctx, domain.ArticleFromRpc(req), views.Article)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) UpdateArticle(ctx context.Context, req *productsRPC.Article) (*emptypb.Empty, error) {
	const op = "grpc.UpdateArticle"
	log.Blue(op)

	_, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return nil, s.s.Update(ctx, domain.ArticleFromRpc(req), views.Article)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) DeleteArticle(ctx context.Context, req *productsRPC.Id) (*emptypb.Empty, error) {
	const op = "grpc.DeleteArticle"
	log.Blue(op)

	_, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return nil, s.s.Delete(ctx, req.GetId(), views.Article)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) GetAllArticles(ctx context.Context, _ *emptypb.Empty) (*productsRPC.ArticleList, error) {
	const op = "grpc.GetAllArticles"
	log.Blue(op)

	res, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return s.s.GetAll(ctx, views.Article)
	})
	if err != nil {
		return nil, err
	}

	articles := res.([]*domain.Article)
	return &productsRPC.ArticleList{
		Items: domain.ArticlesToRpc(articles),
	}, nil
}

func (s *ServerAPI) GetArticle(ctx context.Context, req *productsRPC.Id) (*productsRPC.Article, error) {
	const op = "grpc.GetArticle"
	log.Blue(op)

	res, err := handleCRUDResponse(ctx, op, func() (any, error) {
		return s.s.Get(ctx, req.GetId(), views.Article)
	})
	if err != nil {
		return nil, err
	}

	return domain.ArticleToRpc(res.(*domain.Article)), nil
}
