package repository

import (
	"context"
	"github.com/autumnterror/volha-backend/pkg/views"
	"testing"

	"github.com/autumnterror/breezynotes/pkg/log"
	"github.com/autumnterror/volha-backend/internal/product-service/config"
	"github.com/autumnterror/volha-backend/internal/product-service/domain"
	"github.com/stretchr/testify/assert"
)

func TestGoodProductColorPhotos(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	db, err := NewConnect(config.Test())
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
	}))

	c := &domain.ProductColorPhotos{ProductId: "test-product", ColorId: color.Id, Photos: []string{"photo1", "photo2"}}
	updated := &domain.ProductColorPhotos{ProductId: "test-product", ColorId: color.Id, Photos: []string{"new photo1", "new photo2"}}

	assert.NoError(t, driver.CreateProductColorPhotos(ctx, c))

	all, err := driver.GetAllProductColorPhotos(ctx)
	assert.NoError(t, err)
	log.Println("product color photos after create:", all)

	assert.NoError(t, driver.UpdateProductColorPhotos(ctx, updated))

	ph, err := driver.GetPhotosByProductColor(ctx, c.ProductId, c.ColorId)
	assert.NoError(t, err)
	log.Println("get photos after update:", ph)

	assert.NoError(t, driver.DeleteProductColorPhotos(ctx, c.ProductId, c.ColorId))
}

func TestBadProductColorPhotos(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	db, err := NewConnect(config.Test())
	assert.NoError(t, err)

	tx, err := db.Driver.Begin()
	assert.NoError(t, err)

	driver := Driver{Driver: tx}
	defer t.Cleanup(func() {
		_ = tx.Rollback()
		db.Disconnect()
	})

	pcp := &domain.ProductColorPhotos{ProductId: "product_test_bad", ColorId: "color_test_bad", Photos: []string{"photo1", "photo2"}}
	tx1, err := db.Driver.Begin()
	assert.NoError(t, err)
	{
		driver := Driver{Driver: tx1}
		assert.ErrorIs(t, driver.CreateProductColorPhotos(ctx, pcp), domain.ErrForeignKey)
	}
	_ = tx1.Rollback()

	assert.ErrorIs(
		t,
		driver.UpdateProductColorPhotos(
			ctx,
			&domain.ProductColorPhotos{ProductId: "not exist", ColorId: "not exist", Photos: []string{"photo1", "photo2"}},
		),
		domain.ErrNotFound,
	)

	assert.ErrorIs(t, driver.DeleteProductColorPhotos(ctx, pcp.ProductId, pcp.ColorId), domain.ErrNotFound)

}
