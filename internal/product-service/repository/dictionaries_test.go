package repository

import (
	"context"
	"github.com/autumnterror/volha-backend/internal/product-service/infra/psql"
	"github.com/autumnterror/volha-backend/pkg/views"

	"github.com/autumnterror/breezynotes/pkg/log"

	"github.com/autumnterror/volha-backend/internal/product-service/config"
	"github.com/autumnterror/volha-backend/internal/product-service/domain"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDictionaries(t *testing.T) {
	db, err := psql.NewConnect(config.Test())
	assert.NoError(t, err)

	ctx := context.Background()

	tx, err := db.Driver.Begin()
	assert.NoError(t, err)

	driver := Driver{Driver: tx}
	defer db.Disconnect()
	t.Cleanup(func() {
		_ = tx.Rollback()
	})

	brand := &domain.Brand{Id: "brand_test-dic", Title: "Test Brand"}
	cat := &domain.Category{Id: "cat_test-dic", Title: "Test Cat", Uri: "test-uri"}
	cat2 := &domain.Category{Id: "cat_test2-dic", Title: "Test Cat", Uri: "test-uri"}
	cou := &domain.Country{Id: "country_test-dic", Title: "Test Country", Friendly: "FriendlyName"}
	m := &domain.Material{Id: "mat_test-dic", Title: "Test Material"}
	col := &domain.Color{Id: "color_test-dic", Title: "TestColor", Hex: "#123456"}
	pmax := &domain.ProductId{
		Id:          "pmax-dic",
		Title:       "test",
		Article:     "test",
		Brand:       brand.Id,
		Category:    cat.Id,
		Country:     cou.Id,
		Width:       12,
		Height:      12,
		Depth:       12,
		Materials:   []string{m.Id},
		Colors:      []string{col.Id},
		Photos:      []string{"test"},
		Seems:       nil,
		Price:       12,
		Description: "test",
	}
	pmin := &domain.ProductId{
		Id:          "pmin-dic",
		Title:       "test2",
		Article:     "test2",
		Brand:       brand.Id,
		Category:    cat2.Id,
		Country:     cou.Id,
		Width:       5,
		Height:      5,
		Depth:       5,
		Materials:   []string{m.Id},
		Colors:      []string{col.Id},
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

	d, err := driver.GetDictionaries(ctx, domain.NotByCategory)

	assert.NoError(t, err)
	log.Green("get all dic ", d)
	assert.Equal(t, 5, int(d.MinHeight))

	d, err = driver.GetDictionaries(ctx, cat.Id)
	assert.NoError(t, err)
	log.Green("get dic by category ", d)
	assert.Equal(t, 12, int(d.MinHeight))
}
