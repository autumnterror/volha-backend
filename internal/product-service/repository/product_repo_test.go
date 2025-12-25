package repository

import (
	"context"
	"testing"

	"github.com/autumnterror/volha-backend/internal/product-service/infra/psql"

	"github.com/autumnterror/volha-backend/pkg/views"

	"github.com/autumnterror/breezynotes/pkg/log"

	"github.com/autumnterror/volha-backend/internal/product-service/config"
	"github.com/autumnterror/volha-backend/internal/product-service/domain"

	"github.com/stretchr/testify/assert"
)

func TestProductGood(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	t.Run("product good bad", func(t *testing.T) {
		t.Parallel()

		db, err := psql.NewConnect(config.Test())
		assert.NoError(t, err)

		tx, err := db.Driver.Begin()
		assert.NoError(t, err)

		driver := Driver{Driver: tx}
		defer t.Cleanup(func() {
			_ = tx.Rollback()
			db.Disconnect()
		})

		brand := &domain.Brand{Id: "brand1", Title: "TestBrand"}
		category := &domain.Category{Id: "cat1", Title: "TestCategory", Uri: "test-cat"}
		country := &domain.Country{Id: "country1", Title: "TestCountry", Friendly: "Friendly"}
		material := &domain.Material{Id: "mat1", Title: "Leather"}
		color := &domain.Color{Id: "color1", Title: "Red", Hex: "#FF0000"}
		assert.NoError(t, driver.Create(ctx, brand, views.Brand))
		assert.NoError(t, driver.Create(ctx, category, views.Category))
		assert.NoError(t, driver.Create(ctx, country, views.Country))
		assert.NoError(t, driver.Create(ctx, material, views.Material))
		assert.NoError(t, driver.Create(ctx, color, views.Color))

		assert.NoError(t, driver.CreateProduct(ctx, &domain.ProductId{
			Id:          "test-product",
			Title:       "test",
			Article:     "test",
			Brand:       brand.Id,
			Category:    category.Id,
			Country:     country.Id,
			Width:       1,
			Height:      2,
			Depth:       3,
			Materials:   []string{"mat1"},
			Colors:      []string{"color1"},
			Photos:      []string{"photo1", "photo2"},
			Seems:       []string{},
			Price:       100,
			Description: "test",
			Views:       1000,
			IsFavorite:  true,
		}))

		p, total, err := driver.GetAllProducts(ctx, 0, 10)
		if assert.NoError(t, err) {
			log.Green("products after create", p, " total: ", total)
		}

		brand = &domain.Brand{Id: "brand2", Title: "new TestBrand"}
		category = &domain.Category{Id: "cat2", Title: "new TestCategory", Uri: "new test-cat"}
		country = &domain.Country{Id: "country2", Title: "new TestCountry", Friendly: "new Friendly"}
		material = &domain.Material{Id: "mat2", Title: "new Leather"}
		color = &domain.Color{Id: "color2", Title: "new Red", Hex: "#FF0000"}
		assert.NoError(t, driver.Create(ctx, brand, views.Brand))
		assert.NoError(t, driver.Create(ctx, category, views.Category))
		assert.NoError(t, driver.Create(ctx, country, views.Country))
		assert.NoError(t, driver.Create(ctx, material, views.Material))
		assert.NoError(t, driver.Create(ctx, color, views.Color))

		newViews := 10000

		assert.NoError(t, driver.UpdateProduct(ctx, &domain.ProductId{
			Id:          "test-product",
			Title:       "new test",
			Article:     "new test",
			Brand:       brand.Id,
			Category:    category.Id,
			Country:     country.Id,
			Width:       1,
			Height:      2,
			Depth:       3,
			Materials:   []string{"mat1", "mat2"},
			Colors:      []string{"color1", "color2"},
			Photos:      []string{"new photo1", "new photo2"},
			Seems:       []string{},
			Price:       100,
			Description: "test",
			Views:       int32(newViews),
			IsFavorite:  false,
		}))

		assert.NoError(t, driver.IncrementViews(ctx, "test-product"))

		pr, err := driver.GetProduct(ctx, "test-product")
		if assert.NoError(t, err) {
			log.Green("product after update ", pr)
		}

		assert.Equal(t, int32(newViews+1), pr.Views)

		res, total, err := driver.SearchProducts(ctx, &domain.ProductSearch{
			Id:      "",
			Title:   "new test",
			Article: "",
			Start:   0,
			Finish:  10,
		})
		if assert.NoError(t, err) && assert.NotEqual(t, 0, len(res)) && assert.NotEqual(t, 0, total) {
			log.Green("search by title ", res)
		}
		res, total, err = driver.SearchProducts(ctx, &domain.ProductSearch{
			Id:      "",
			Title:   "",
			Article: "new test",
			Start:   0,
			Finish:  0,
		})
		if assert.NoError(t, err) && assert.Equal(t, 0, len(res)) && assert.NotEqual(t, 0, total) {
			log.Green("search by article full ", res)
		}
		res, total, err = driver.SearchProducts(ctx, &domain.ProductSearch{
			Id:      "",
			Title:   "ew",
			Article: "",
			Start:   0,
			Finish:  10,
		})
		if assert.NoError(t, err) && assert.NotEqual(t, 0, len(res)) && assert.NotEqual(t, 0, total) {
			log.Green("search by article not full ", res)
		}

		assert.NoError(t, driver.DeleteProduct(ctx, "test-product"))
	})
}

func TestProductFilters(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	db, err := psql.NewConnect(config.Test())
	assert.NoError(t, err)

	tx, err := db.Driver.Begin()
	assert.NoError(t, err)

	driver := Driver{Driver: tx}
	defer db.Disconnect()
	t.Cleanup(func() { _ = tx.Rollback() })

	brand := &domain.Brand{Id: "brand1-filter", Title: "TestBrand"}
	category := &domain.Category{Id: "cat1-filter", Title: "TestCategory", Uri: "test-cat"}
	country := &domain.Country{Id: "country1-filter", Title: "TestCountry", Friendly: "Friendly"}
	material := &domain.Material{Id: "mat1-filter", Title: "Leather"}
	color := &domain.Color{Id: "color1-filter", Title: "Red", Hex: "#FF0000"}
	assert.NoError(t, driver.Create(ctx, brand, views.Brand))
	assert.NoError(t, driver.Create(ctx, category, views.Category))
	assert.NoError(t, driver.Create(ctx, country, views.Country))
	assert.NoError(t, driver.Create(ctx, material, views.Material))
	assert.NoError(t, driver.Create(ctx, color, views.Color))

	product := &domain.ProductId{
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
		Views:       1000,
		IsFavorite:  true,
	}
	assert.NoError(t, driver.CreateProduct(ctx, product))

	tests := []struct {
		name   string
		filter domain.ProductFilter
	}{
		{"by brand", domain.ProductFilter{Brand: []string{brand.Id}}},
		{"by category", domain.ProductFilter{Category: []string{category.Id}}},
		{"by country", domain.ProductFilter{Country: []string{country.Id}}},
		{"by material", domain.ProductFilter{Materials: []string{material.Id}}},
		{"by color", domain.ProductFilter{Colors: []string{color.Id}}},
		{"by size", domain.ProductFilter{MinWidth: 50, MaxWidth: 70, MinHeight: 100, MaxHeight: 120, MinDepth: 30, MaxDepth: 50}},
		{"by price", domain.ProductFilter{MinPrice: 500, MaxPrice: 1000}},
		{"with sort asc", domain.ProductFilter{SortBy: "price", SortOrder: "asc"}},
		{"with sort desc", domain.ProductFilter{SortBy: "width", SortOrder: "desc"}},
		{"with sort asc", domain.ProductFilter{SortBy: "views", SortOrder: "asc"}},
		{"with limit and offset", domain.ProductFilter{Limit: 1, Offset: 0}},
		{"is_favorite", domain.ProductFilter{IsFavorite: true}},
		{"by title", domain.ProductFilter{Title: "terPro"}},
		{"complex", domain.ProductFilter{
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
			result, total, err := driver.FilterProducts(ctx, &tt.filter)
			assert.NoError(t, err)
			assert.NotEmpty(t, result)
			assert.NotEqual(t, 0, total)
			log.Printf("Filter '%s' returned %d product(s)\n", tt.name, len(result))
		})
	}
}
