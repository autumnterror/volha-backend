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

func TestGoodSimpleTables(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	t.Run("brands", func(t *testing.T) {
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

		brand := &productsRPC.Brand{Id: "brand_test", Title: "Test Brand"}
		updated := &productsRPC.Brand{Id: "brand_test", Title: "Updated Brand"}

		assert.NoError(t, driver.Create(ctx, brand, views.Brand))

		all, err := driver.GetAll(ctx, views.Brand)
		assert.NoError(t, err)
		log.Println("brands after create:", all)

		assert.NoError(t, driver.Update(ctx, updated, views.Brand))

		all, err = driver.GetAll(ctx, views.Brand)
		assert.NoError(t, err)
		log.Println("brands after update:", all)

		assert.NoError(t, driver.Delete(ctx, brand.Id, views.Brand))
	})

	t.Run("categories", func(t *testing.T) {
		t.Parallel()

		db, err := NewConnect(config.Test())
		assert.NoError(t, err)

		tx, err := db.Driver.Begin()
		assert.NoError(t, err)

		driver := Driver{Driver: tx}

		t.Cleanup(func() {
			_ = tx.Rollback()
			db.Disconnect()
		})

		c := &productsRPC.Category{Id: "cat_test", Title: "Test Cat", Uri: "test-uri"}
		updated := &productsRPC.Category{Id: "cat_test", Title: "Updated", Uri: "upd-uri"}

		assert.NoError(t, driver.Create(ctx, c, views.Category))

		all, err := driver.GetAll(ctx, views.Category)
		assert.NoError(t, err)
		log.Println("categories after create:", all)

		assert.NoError(t, driver.Update(ctx, updated, views.Category))

		all, err = driver.GetAll(ctx, views.Category)
		assert.NoError(t, err)
		log.Println("categories after update:", all)

		assert.NoError(t, driver.Delete(ctx, c.Id, views.Category))
	})

	t.Run("countries", func(t *testing.T) {
		t.Parallel()

		db, err := NewConnect(config.Test())
		assert.NoError(t, err)

		tx, err := db.Driver.Begin()
		assert.NoError(t, err)

		driver := Driver{Driver: tx}

		t.Cleanup(func() {
			_ = tx.Rollback()
			db.Disconnect()
		})

		c := &productsRPC.Country{Id: "country_test", Title: "Test Country", Friendly: "FriendlyName"}
		updated := &productsRPC.Country{Id: "country_test", Title: "Updated", Friendly: "UpdatedFriendly"}

		assert.NoError(t, driver.Create(ctx, c, views.Country))

		all, err := driver.GetAll(ctx, views.Country)
		assert.NoError(t, err)
		log.Println("countries after create:", all)

		assert.NoError(t, driver.Update(ctx, updated, views.Country))

		all, err = driver.GetAll(ctx, views.Country)
		assert.NoError(t, err)
		log.Println("countries after update:", all)

		assert.NoError(t, driver.Delete(ctx, c.Id, views.Country))
	})

	t.Run("materials", func(t *testing.T) {
		t.Parallel()

		db, err := NewConnect(config.Test())
		assert.NoError(t, err)

		tx, err := db.Driver.Begin()
		assert.NoError(t, err)

		driver := Driver{Driver: tx}
		defer db.Disconnect()
		t.Cleanup(func() { _ = tx.Rollback() })

		m := &productsRPC.Material{Id: "mat_test", Title: "Test Material"}
		updated := &productsRPC.Material{Id: "mat_test", Title: "Updated Material"}

		assert.NoError(t, driver.Create(ctx, m, views.Material))

		all, err := driver.GetAll(ctx, views.Material)
		assert.NoError(t, err)
		log.Println("materials after create:", all)

		assert.NoError(t, driver.Update(ctx, updated, views.Material))

		all, err = driver.GetAll(ctx, views.Material)
		assert.NoError(t, err)
		log.Println("materials after update:", all)

		assert.NoError(t, driver.Delete(ctx, m.Id, views.Material))
	})

	t.Run("colors", func(t *testing.T) {
		t.Parallel()

		db, err := NewConnect(config.Test())
		assert.NoError(t, err)

		tx, err := db.Driver.Begin()
		assert.NoError(t, err)

		driver := Driver{Driver: tx}

		t.Cleanup(func() {
			_ = tx.Rollback()
			db.Disconnect()
		})

		c := &productsRPC.Color{Id: "color_test", Title: "TestColor", Hex: "#123456"}
		updated := &productsRPC.Color{Id: "color_test", Title: "UpdatedColor", Hex: "#654321"}

		assert.NoError(t, driver.Create(ctx, c, views.Color))

		all, err := driver.GetAll(ctx, views.Color)
		assert.NoError(t, err)
		log.Println("colors after create:", all)

		assert.NoError(t, driver.Update(ctx, updated, views.Color))

		all, err = driver.GetAll(ctx, views.Color)
		assert.NoError(t, err)
		log.Println("colors after update:", all)

		assert.NoError(t, driver.Delete(ctx, c.Id, views.Color))
	})
}

func TestBadSimpleTables(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	t.Run("brands bad", func(t *testing.T) {
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

		assert.ErrorIs(t, driver.Create(ctx, 2, views.Brand), ErrInvalidType)
		assert.ErrorIs(t, driver.Create(ctx, 2, 19), ErrUnknownType)

		brand := &productsRPC.Brand{Id: "brand_test_bad", Title: "Test Brand"}
		tx1, err := db.Driver.Begin()
		assert.NoError(t, err)
		{
			driver := Driver{Driver: tx1}
			assert.NoError(t, driver.Create(ctx, brand, views.Brand))
			assert.ErrorIs(t, driver.Create(ctx, brand, views.Brand), ErrAlreadyExists)
		}
		_ = tx1.Rollback()

		assert.ErrorIs(
			t,
			driver.Update(
				ctx,
				&productsRPC.Brand{Id: "not_exist", Title: "Test Brand"},
				views.Brand,
			),
			ErrNotFound,
		)

		assert.ErrorIs(t, driver.Delete(ctx, brand.Id, views.Brand), ErrNotFound)
	})

	t.Run("categories bad", func(t *testing.T) {
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

		assert.ErrorIs(t, driver.Create(ctx, 2, views.Category), ErrInvalidType)
		assert.ErrorIs(t, driver.Create(ctx, 2, 19), ErrUnknownType)

		cat := &productsRPC.Category{Id: "cat_test_bad", Title: "Test Cat", Uri: "test-uri"}

		tx1, err := db.Driver.Begin()
		assert.NoError(t, err)
		{
			driver := Driver{Driver: tx1}
			assert.NoError(t, driver.Create(ctx, cat, views.Category))
			assert.ErrorIs(t, driver.Create(ctx, cat, views.Category), ErrAlreadyExists)
		}
		_ = tx1.Rollback()

		assert.ErrorIs(
			t,
			driver.Update(
				ctx,
				&productsRPC.Category{Id: "not_exist", Title: "Test Cat", Uri: "test-uri"},
				views.Category,
			),
			ErrNotFound,
		)

		assert.ErrorIs(t, driver.Delete(ctx, cat.Id, views.Category), ErrNotFound)
	})

	t.Run("countries bad", func(t *testing.T) {
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

		assert.ErrorIs(t, driver.Create(ctx, 2, views.Country), ErrInvalidType)
		assert.ErrorIs(t, driver.Create(ctx, 2, 19), ErrUnknownType)

		country := &productsRPC.Country{Id: "country_test_bad", Title: "Test Country", Friendly: "FriendlyName"}

		tx1, err := db.Driver.Begin()
		assert.NoError(t, err)
		{
			driver := Driver{Driver: tx1}
			assert.NoError(t, driver.Create(ctx, country, views.Country))
			assert.ErrorIs(t, driver.Create(ctx, country, views.Country), ErrAlreadyExists)
		}
		_ = tx1.Rollback()

		assert.ErrorIs(
			t,
			driver.Update(
				ctx,
				&productsRPC.Country{Id: "not_exist", Title: "Test Country", Friendly: "FriendlyName"},
				views.Country,
			),
			ErrNotFound,
		)

		assert.ErrorIs(t, driver.Delete(ctx, country.Id, views.Country), ErrNotFound)
	})

	t.Run("materials bad", func(t *testing.T) {
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

		assert.ErrorIs(t, driver.Create(ctx, 2, views.Material), ErrInvalidType)
		assert.ErrorIs(t, driver.Create(ctx, 2, 19), ErrUnknownType)

		mat := &productsRPC.Material{Id: "mat_test_bad", Title: "Test Material"}
		tx1, err := db.Driver.Begin()
		assert.NoError(t, err)
		{
			driver := Driver{Driver: tx1}
			assert.NoError(t, driver.Create(ctx, mat, views.Material))
			assert.ErrorIs(t, driver.Create(ctx, mat, views.Material), ErrAlreadyExists)
		}
		_ = tx1.Rollback()

		assert.ErrorIs(
			t,
			driver.Update(
				ctx,
				&productsRPC.Material{Id: "not_exist", Title: "Test Material"},
				views.Material,
			),
			ErrNotFound,
		)

		assert.ErrorIs(t, driver.Delete(ctx, mat.Id, views.Material), ErrNotFound)
	})

	t.Run("colors bad", func(t *testing.T) {
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

		assert.ErrorIs(t, driver.Create(ctx, 2, views.Color), ErrInvalidType)
		assert.ErrorIs(t, driver.Create(ctx, 2, 19), ErrUnknownType)

		color := &productsRPC.Color{Id: "color_test_bad", Title: "TestColor", Hex: "#123456"}
		tx1, err := db.Driver.Begin()
		assert.NoError(t, err)
		{
			driver := Driver{Driver: tx1}
			assert.NoError(t, driver.Create(ctx, color, views.Color))
			assert.ErrorIs(t, driver.Create(ctx, color, views.Color), ErrAlreadyExists)
		}
		_ = tx1.Rollback()

		assert.ErrorIs(
			t,
			driver.Update(
				ctx,
				&productsRPC.Color{Id: "not_exist", Title: "TestColor", Hex: "#123456"},
				views.Color,
			),
			ErrNotFound,
		)

		assert.ErrorIs(t, driver.Delete(ctx, color.Id, views.Color), ErrNotFound)
	})
}
