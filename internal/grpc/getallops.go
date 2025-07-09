package grpc

import (
	"context"
	productsRPC "github.com/autumnterror/volha-proto/gen/products"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"productService/internal/utils/format"
	"productService/internal/views"
)

func (s *ServerAPI) FilterProducts(
	ctx context.Context,
	req *productsRPC.ProductFilter,
) (*productsRPC.ProductList, error) {
	const op = "productsRPC.ServerAPI.FilterProducts"

	type result struct {
		lp  *productsRPC.ProductList
		err error
	}
	res := make(chan result, 1)

	go func() {
		filter := convertToProductFilter(req)

		lp, err := s.API.FilterProducts(filter)
		if err != nil {
			res <- result{err: format.Error(op, status.Error(codes.Internal, err.Error()))}
			log.Println(format.Error(op, status.Error(codes.Internal, err.Error())))
			return
		}

		res <- result{lp: convertToProductList(lp), err: nil}
		log.Printf("Get all with filter request. Filter %s\n", format.Struct(filter))

	}()

	select {
	case <-ctx.Done():
		log.Println(format.Error(op, status.Error(codes.DeadlineExceeded, "Context dead")))
		return nil, format.Error(op, status.Error(codes.DeadlineExceeded, "Context dead"))
	case r := <-res:
		return r.lp, r.err
	}
}

func (s *ServerAPI) GetAll(
	ctx context.Context,
	req *emptypb.Empty,
) (*productsRPC.ProductList, error) {
	const op = "productsRPC.ServerAPI.GetAll"

	type result struct {
		lp  *productsRPC.ProductList
		err error
	}
	res := make(chan result, 1)

	go func() {
		lp, err := s.API.GetAll()
		if err != nil {
			res <- result{err: format.Error(op, status.Error(codes.Internal, err.Error()))}
			log.Println(format.Error(op, status.Error(codes.Internal, err.Error())))
			return
		}

		res <- result{lp: convertToProductList(lp), err: nil}

		log.Printf("Get all request.\n")

	}()

	select {
	case <-ctx.Done():
		log.Println(format.Error(op, status.Error(codes.DeadlineExceeded, "Context dead")))
		return nil, format.Error(op, status.Error(codes.DeadlineExceeded, "Context dead"))
	case r := <-res:
		return r.lp, r.err
	}
}

func convertToProductList(viewProducts []views.Product) *productsRPC.ProductList {
	productList := &productsRPC.ProductList{}
	for _, p := range viewProducts {
		product := &productsRPC.Product{
			Id:          p.Id,
			Title:       p.Title,
			Article:     p.Article,
			Brand:       p.Brand,
			Country:     p.Country,
			Width:       int32(p.Width),
			Height:      int32(p.Height),
			Depth:       int32(p.Depth),
			Materials:   p.Materials,
			Color:       p.Colors,
			Photos:      p.Photos,
			Seems:       p.Seems,
			Price:       int32(p.Price),
			Description: p.Description,
		}
		productList.Products = append(productList.Products, product)
	}
	return productList
}
func convertToProductFilter(p *productsRPC.ProductFilter) *views.ProductFilter {
	return &views.ProductFilter{
		Brand:     p.Brand,
		Country:   p.Country,
		MinWidth:  int(p.MinWidth),
		MaxWidth:  int(p.MaxWidth),
		MinHeight: int(p.MinHeight),
		MaxHeight: int(p.MaxHeight),
		MinDepth:  int(p.MinDepth),
		MaxDepth:  int(p.MaxDepth),
		Materials: p.Materials,
		Colors:    p.Colors,
		MinPrice:  int(p.MinPrice),
		MaxPrice:  int(p.MaxPrice),
		SortBy:    p.SortBy,
		SortOrder: p.SortOrder,
		Offset:    int(p.Offset),
		Limit:     int(p.Limit),
	}
}
