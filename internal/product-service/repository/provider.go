package repository

import "context"

type Provider interface {
	Product(ctx context.Context) ProductRepo
	Basic(ctx context.Context) BasicRepo
	Dict(ctx context.Context) DictRepo
	Pcp(ctx context.Context) ProductColorPhotosRepo
}
