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

func TestGetDictionaries(t *testing.T) {
	db, err := NewConnect(config.Test())
	assert.NoError(t, err)

	ctx := context.Background()

	tx, err := db.Driver.Begin()
	assert.NoError(t, err)

	driver := Driver{Driver: tx}
	defer db.Disconnect()
	t.Cleanup(func() {
		_ = tx.Rollback()
	})

	brand := &productsRPC.Brand{Id: "brand_test-dic", Title: "Test Brand"}
	cat := &productsRPC.Category{Id: "cat_test-dic", Title: "Test Cat", Uri: "test-uri"}
	cat2 := &productsRPC.Category{Id: "cat_test2-dic", Title: "Test Cat", Uri: "test-uri"}
	cou := &productsRPC.Country{Id: "country_test-dic", Title: "Test Country", Friendly: "FriendlyName"}
	m := &productsRPC.Material{Id: "mat_test-dic", Title: "Test Material"}
	col := &productsRPC.Color{Id: "color_test-dic", Title: "TestColor", Hex: "#123456"}
	pmax := &productsRPC.ProductId{
		Id:          "pmax-dic",
		Title:       "test",
		Article:     "test",
		Brand:       brand.GetId(),
		Category:    cat.Id,
		Country:     cou.Id,
		Width:       12,
		Height:      12,
		Depth:       12,
		Materials:   []string{m.GetId()},
		Colors:      []string{col.GetId()},
		Photos:      []string{"test"},
		Seems:       nil,
		Price:       12,
		Description: "test",
	}
	pmin := &productsRPC.ProductId{
		Id:          "pmin-dic",
		Title:       "test2",
		Article:     "test2",
		Brand:       brand.GetId(),
		Category:    cat2.Id,
		Country:     cou.Id,
		Width:       5,
		Height:      5,
		Depth:       5,
		Materials:   []string{m.GetId()},
		Colors:      []string{col.GetId()},
		Photos:      []string{"test"},
		Seems:       []string{"test-pmax"},
		Price:       5,
		Description: "test",
	}
	err = driver.Create(ctx, brand, views.Brand)
	assert.NoError(t, err)
	err = driver.Create(ctx, cat, views.Category)
	assert.NoError(t, err)
	err = driver.Create(ctx, cat2, views.Category)
	assert.NoError(t, err)
	err = driver.Create(ctx, cou, views.Country)
	assert.NoError(t, err)
	err = driver.Create(ctx, m, views.Material)
	assert.NoError(t, err)
	err = driver.Create(ctx, col, views.Color)
	assert.NoError(t, err)
	assert.NoError(t, driver.CreateProduct(ctx, pmax))
	assert.NoError(t, driver.CreateProduct(ctx, pmin))

	d, err := driver.GetDictionaries(ctx, NotByCategory)

	assert.NoError(t, err)
	log.Green("get all dic ", d)
	assert.Equal(t, 5, int(d.MinHeight))

	d, err = driver.GetDictionaries(ctx, cat.GetId())
	assert.NoError(t, err)
	log.Green("get dic by category ", d)
	assert.Equal(t, 12, int(d.MinHeight))
}
