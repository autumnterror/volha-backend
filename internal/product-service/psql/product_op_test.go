package psql

import (
	"context"
	"github.com/autumnterror/breezynotes/pkg/log"
	"testing"

	"github.com/autumnterror/volha-backend/internal/product-service/config"
	productsRPC "github.com/autumnterror/volha-backend/pkg/proto/gen"
	"github.com/autumnterror/volha-backend/pkg/views"
	"github.com/stretchr/testify/assert"
)

func TestProductGood(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	t.Run("product good bad", func(t *testing.T) {
		t.Parallel()

		db, err := NewConnect(config.Test())
		assert.NoError(t, err)

		tx, err := db.Driver.Begin()
		assert.NoError(t, err)

		driver := Driver{Driver: tx}
		defer t.Cleanup(func() {
			_ = tx.Rollback()
			db.Disconnect()
		})

		brand := &productsRPC.Brand{Id: "brand1", Title: "TestBrand"}
		category := &productsRPC.Category{Id: "cat1", Title: "TestCategory", Uri: "test-cat"}
		country := &productsRPC.Country{Id: "country1", Title: "TestCountry", Friendly: "Friendly"}
		material := &productsRPC.Material{Id: "mat1", Title: "Leather"}
		color := &productsRPC.Color{Id: "color1", Title: "Red", Hex: "#FF0000"}
		assert.NoError(t, driver.Create(ctx, brand, views.Brand))
		assert.NoError(t, driver.Create(ctx, category, views.Category))
		assert.NoError(t, driver.Create(ctx, country, views.Country))
		assert.NoError(t, driver.Create(ctx, material, views.Material))
		assert.NoError(t, driver.Create(ctx, color, views.Color))

		assert.NoError(t, driver.CreateProduct(ctx, &productsRPC.ProductId{
			Id:          "test-product",
			Title:       "test",
			Article:     "test",
			Brand:       brand.GetId(),
			Category:    category.GetId(),
			Country:     country.GetId(),
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

		p, err := driver.GetAllProducts(ctx, 0, 10)
		if assert.NoError(t, err) {
			log.Green("products after create", p)
		}

		brand = &productsRPC.Brand{Id: "brand2", Title: "new TestBrand"}
		category = &productsRPC.Category{Id: "cat2", Title: "new TestCategory", Uri: "new test-cat"}
		country = &productsRPC.Country{Id: "country2", Title: "new TestCountry", Friendly: "new Friendly"}
		material = &productsRPC.Material{Id: "mat2", Title: "new Leather"}
		color = &productsRPC.Color{Id: "color2", Title: "new Red", Hex: "#FF0000"}
		assert.NoError(t, driver.Create(ctx, brand, views.Brand))
		assert.NoError(t, driver.Create(ctx, category, views.Category))
		assert.NoError(t, driver.Create(ctx, country, views.Country))
		assert.NoError(t, driver.Create(ctx, material, views.Material))
		assert.NoError(t, driver.Create(ctx, color, views.Color))

		assert.NoError(t, driver.UpdateProduct(ctx, &productsRPC.ProductId{
			Id:          "test-product",
			Title:       "new test",
			Article:     "new test",
			Brand:       brand.GetId(),
			Category:    category.GetId(),
			Country:     country.GetId(),
			Width:       1,
			Height:      2,
			Depth:       3,
			Materials:   []string{"mat1", "mat2"},
			Colors:      []string{"color1", "color2"},
			Photos:      []string{"new photo1", "new photo2"},
			Seems:       []string{},
			Price:       100,
			Description: "test",
		}))

		pr, err := driver.GetProduct(ctx, "test-product")
		if assert.NoError(t, err) {
			log.Green("product after update ", pr)
		}

		res, err := driver.SearchProducts(ctx, &productsRPC.ProductSearch{
			Id:      "",
			Title:   "new test",
			Article: "",
		})
		if assert.NoError(t, err) && assert.NotEqual(t, 0, len(res)) {
			log.Green("search by title ", res)
		}
		res, err = driver.SearchProducts(ctx, &productsRPC.ProductSearch{
			Id:      "",
			Title:   "",
			Article: "new test",
		})
		if assert.NoError(t, err) && assert.NotEqual(t, 0, len(res)) {
			log.Green("search by article full ", res)
		}
		res, err = driver.SearchProducts(ctx, &productsRPC.ProductSearch{
			Id:      "",
			Title:   "ew",
			Article: "",
		})
		if assert.NoError(t, err) && assert.NotEqual(t, 0, len(res)) {
			log.Green("search by article not full ", res)
		}

		assert.NoError(t, driver.DeleteProduct(ctx, "test-product"))
	})

}
