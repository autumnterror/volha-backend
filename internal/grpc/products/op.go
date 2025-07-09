package products

import (
	"context"
	"gateway/internal/utils/convert"
	"gateway/internal/utils/format"
	"gateway/internal/views"
	productsRPC "github.com/autumnterror/volha-proto/gen/products"
)

func (c *Client) GetAllFilter(ctx context.Context, pf *views.ProductFilter) ([]views.Product, error) {
	const op = "grpc.products.GetAllFilter"

	lp, err := c.api.FilterProducts(ctx, convert.ToProductFilter(pf))
	if err != nil {
		return nil, format.Error(op, err)
	}

	return convert.ToProductViewList(lp), nil
}

func (c *Client) GetAll(ctx context.Context) ([]views.Product, error) {
	const op = "grpc.client.GetAll"

	lp, err := c.api.GetAll(ctx, nil)
	if err != nil {
		return nil, format.Error(op, err)
	}

	return convert.ToProductViewList(lp), nil
}

func (c *Client) Create(ctx context.Context, p *views.Product) error {
	const op = "grpc.client.Create"

	if _, err := c.api.Create(ctx, convert.ToProductRPC(p)); err != nil {
		return format.Error(op, err)
	}

	return nil
}

func (c *Client) Update(ctx context.Context, p *views.Product) error {
	const op = "grpc.client.Update"

	if _, err := c.api.Update(ctx, convert.ToProductRPC(p)); err != nil {
		return format.Error(op, err)
	}

	return nil
}

func (c *Client) Delete(ctx context.Context, id string) error {
	const op = "grpc.client.Delete"

	if _, err := c.api.Delete(ctx, &productsRPC.Id{Id: id}); err != nil {
		return format.Error(op, err)
	}

	return nil
}
