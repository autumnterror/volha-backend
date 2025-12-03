package psql

import (
	"context"
	"github.com/autumnterror/breezynotes/pkg/log"
	"github.com/autumnterror/volha-backend/internal/product-service/config"
	productsRPC "github.com/autumnterror/volha-backend/pkg/proto/gen"
	"github.com/autumnterror/volha-backend/pkg/views"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductFilters(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	db, err := NewConnect(config.Test())
	assert.NoError(t, err)

	tx, err := db.Driver.Begin()
	assert.NoError(t, err)

	driver := Driver{Driver: tx}
	defer db.Disconnect()
	t.Cleanup(func() { _ = tx.Rollback() })

	brand := &productsRPC.Brand{Id: "brand1-filter", Title: "TestBrand"}
	category := &productsRPC.Category{Id: "cat1-filter", Title: "TestCategory", Uri: "test-cat"}
	country := &productsRPC.Country{Id: "country1-filter", Title: "TestCountry", Friendly: "Friendly"}
	material := &productsRPC.Material{Id: "mat1-filter", Title: "Leather"}
	color := &productsRPC.Color{Id: "color1-filter", Title: "Red", Hex: "#FF0000"}
	assert.NoError(t, driver.Create(ctx, brand, views.Brand))
	assert.NoError(t, driver.Create(ctx, category, views.Category))
	assert.NoError(t, driver.Create(ctx, country, views.Country))
	assert.NoError(t, driver.Create(ctx, material, views.Material))
	assert.NoError(t, driver.Create(ctx, color, views.Color))

	product := &productsRPC.ProductId{
		Id:          "prod1",
		Title:       "FilterProduct",
		Article:     "F-001",
		Brand:       brand.Id,
		Category:    category.Id,
		Country:     country.Id,
		Width:       60,
		Height:      110,
		Depth:       40,
		Materials:   []string{material.Id},
		Colors:      []string{color.Id},
		Photos:      []string{"x.jpg"},
		Seems:       []string{},
		Price:       888,
		Description: "Filtered",
	}
	assert.NoError(t, driver.CreateProduct(ctx, product))

	tests := []struct {
		name   string
		filter productsRPC.ProductFilter
	}{
		{"by brand", productsRPC.ProductFilter{Brand: []string{brand.Id}}},
		{"by category", productsRPC.ProductFilter{Category: []string{category.Id}}},
		{"by country", productsRPC.ProductFilter{Country: []string{country.Id}}},
		{"by material", productsRPC.ProductFilter{Materials: []string{material.Id}}},
		{"by color", productsRPC.ProductFilter{Colors: []string{color.Id}}},
		{"by size", productsRPC.ProductFilter{MinWidth: 50, MaxWidth: 70, MinHeight: 100, MaxHeight: 120, MinDepth: 30, MaxDepth: 50}},
		{"by price", productsRPC.ProductFilter{MinPrice: 500, MaxPrice: 1000}},
		{"with sort asc", productsRPC.ProductFilter{SortBy: "price", SortOrder: "asc"}},
		{"with sort desc", productsRPC.ProductFilter{SortBy: "width", SortOrder: "desc"}},
		{"with limit and offset", productsRPC.ProductFilter{Limit: 1, Offset: 0}},
		{"complex", productsRPC.ProductFilter{
			Brand:     []string{brand.Id},
			Category:  []string{category.Id},
			Country:   []string{country.Id},
			Materials: []string{material.Id},
			Colors:    []string{color.Id},
			MinWidth:  50, MaxWidth: 70,
			MinHeight: 100, MaxHeight: 120,
			MinDepth: 30, MaxDepth: 50,
			MinPrice: 800, MaxPrice: 1000,
			SortBy: "title", SortOrder: "asc",
			Limit: 10,
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := driver.FilterProducts(ctx, &tt.filter)
			assert.NoError(t, err)
			assert.NotEmpty(t, result)
			log.Printf("Filter '%s' returned %d product(s)\n", tt.name, len(result))
		})
	}
}
