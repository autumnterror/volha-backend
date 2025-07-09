package convert

import (
	"gateway/internal/views"
	productsRPC "github.com/autumnterror/volha-proto/gen/products"
)

func ToProductViewList(pds *productsRPC.ProductList) []views.Product {
	var productList []views.Product

	for _, p := range pds.Products {
		product := ToProductView(p)
		productList = append(productList, *product)
	}
	return productList
}

func ToProductList(viewProducts []views.Product) *productsRPC.ProductList {
	productList := &productsRPC.ProductList{}
	for _, p := range viewProducts {
		product := ToProductRPC(&p)
		productList.Products = append(productList.Products, product)
	}
	return productList
}

func ToProductFilter(p *views.ProductFilter) *productsRPC.ProductFilter {
	return &productsRPC.ProductFilter{
		Brand:     p.Brand,
		Country:   p.Country,
		MinWidth:  int32(p.MinWidth),
		MaxWidth:  int32(p.MaxWidth),
		MinHeight: int32(p.MinHeight),
		MaxHeight: int32(p.MaxHeight),
		MinDepth:  int32(p.MinDepth),
		MaxDepth:  int32(p.MaxDepth),
		Materials: p.Materials,
		Colors:    p.Colors,
		MinPrice:  int32(p.MinPrice),
		MaxPrice:  int32(p.MaxPrice),
		SortBy:    p.SortBy,
		SortOrder: p.SortOrder,
		Offset:    int32(p.Offset),
		Limit:     int32(p.Limit),
	}
}

func ToProductFilterView(p *productsRPC.ProductFilter) *views.ProductFilter {
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

func ToProductView(p *productsRPC.Product) *views.Product {
	return &views.Product{
		Id:          p.Id,
		Title:       p.Title,
		Article:     p.Article,
		Brand:       p.Brand,
		Country:     p.Country,
		Width:       int(p.Width),
		Height:      int(p.Height),
		Depth:       int(p.Depth),
		Materials:   p.Materials,
		Colors:      p.Color,
		Photos:      p.Photos,
		Seems:       p.Seems,
		Price:       int(p.Price),
		Description: p.Description,
	}
}

func ToProductRPC(p *views.Product) *productsRPC.Product {
	return &productsRPC.Product{
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
}
