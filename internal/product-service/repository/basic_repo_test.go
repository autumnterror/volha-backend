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

		brand := &domain.Brand{Id: "brand_test", Title: "Test Brand"}
		updated := &domain.Brand{Id: "brand_test", Title: "Updated Brand"}

		assert.NoError(t, driver.Create(ctx, brand, views.Brand))

		all, err := driver.GetAll(ctx, views.Brand)
		assert.NoError(t, err)
		log.Println("brands after create:", all)

		assert.NoError(t, driver.Update(ctx, updated, views.Brand))

		all, err = driver.Get(ctx, brand.Id, views.Brand)
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

		c := &domain.Category{Id: "cat_test", Title: "Test Cat", Uri: "test-uri"}
		updated := &domain.Category{Id: "cat_test", Title: "Updated", Uri: "upd-uri"}

		assert.NoError(t, driver.Create(ctx, c, views.Category))

		all, err := driver.GetAll(ctx, views.Category)
		assert.NoError(t, err)
		log.Println("categories after create:", all)

		assert.NoError(t, driver.Update(ctx, updated, views.Category))

		all, err = driver.Get(ctx, c.Id, views.Category)
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

		c := &domain.Country{Id: "country_test", Title: "Test Country", Friendly: "FriendlyName"}
		updated := &domain.Country{Id: "country_test", Title: "Updated", Friendly: "UpdatedFriendly"}

		assert.NoError(t, driver.Create(ctx, c, views.Country))

		all, err := driver.GetAll(ctx, views.Country)
		assert.NoError(t, err)
		log.Println("countries after create:", all)

		assert.NoError(t, driver.Update(ctx, updated, views.Country))

		all, err = driver.Get(ctx, c.Id, views.Country)
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

		m := &domain.Material{Id: "mat_test", Title: "Test Material"}
		updated := &domain.Material{Id: "mat_test", Title: "Updated Material"}

		assert.NoError(t, driver.Create(ctx, m, views.Material))

		all, err := driver.GetAll(ctx, views.Material)
		assert.NoError(t, err)
		log.Println("materials after create:", all)

		assert.NoError(t, driver.Update(ctx, updated, views.Material))

		all, err = driver.Get(ctx, m.Id, views.Material)
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

		c := &domain.Color{Id: "color_test", Title: "TestColor", Hex: "#123456"}
		updated := &domain.Color{Id: "color_test", Title: "UpdatedColor", Hex: "#654321"}

		assert.NoError(t, driver.Create(ctx, c, views.Color))

		all, err := driver.GetAll(ctx, views.Color)
		assert.NoError(t, err)
		log.Println("colors after create:", all)

		assert.NoError(t, driver.Update(ctx, updated, views.Color))

		all, err = driver.Get(ctx, c.Id, views.Color)
		assert.NoError(t, err)
		log.Println("colors after update:", all)

		assert.NoError(t, driver.Delete(ctx, c.Id, views.Color))
	})

	t.Run("slide", func(t *testing.T) {
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

		s := &domain.Slide{Id: "slide test", Link: "http://example.com", Img: "img1.jpg", Img762: "img762.jpg"}
		updated := &domain.Slide{Id: "slide test", Link: "http://new.example.com", Img: "newimg1.jpg", Img762: "newimg762.jpg"}

		assert.NoError(t, driver.Create(ctx, s, views.Slide))

		all, err := driver.GetAll(ctx, views.Slide)
		assert.NoError(t, err)
		log.Println("Slides after create:", all)

		assert.NoError(t, driver.Update(ctx, updated, views.Slide))

		all, err = driver.Get(ctx, s.Id, views.Slide)
		assert.NoError(t, err)
		log.Println("Slides after update:", all)

		assert.NoError(t, driver.Delete(ctx, s.Id, views.Slide))
	})

	t.Run("article", func(t *testing.T) {
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

		s := &domain.Article{Id: "article test", Title: "Title", Img: "img.jpg", Text: "# zagolovok"}
		updated := &domain.Article{Id: "article test", Title: "newTitle", Img: "newimg.jpg", Text: "# newzagolovok"}

		assert.NoError(t, driver.Create(ctx, s, views.Article))

		all, err := driver.GetAll(ctx, views.Article)
		assert.NoError(t, err)
		log.Println("Articles after create:", all)

		assert.NoError(t, driver.Update(ctx, updated, views.Article))

		all, err = driver.Get(ctx, s.Id, views.Article)
		assert.NoError(t, err)
		log.Println("Articles after update:", all)

		assert.NoError(t, driver.Delete(ctx, s.Id, views.Article))
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

		assert.ErrorIs(t, driver.Create(ctx, 2, views.Brand), domain.ErrInvalidType)
		assert.ErrorIs(t, driver.Create(ctx, 2, 19), domain.ErrUnknownType)

		brand := &domain.Brand{Id: "brand_test_bad", Title: "Test Brand"}
		tx1, err := db.Driver.Begin()
		assert.NoError(t, err)
		{
			driver := Driver{Driver: tx1}
			assert.NoError(t, driver.Create(ctx, brand, views.Brand))
			assert.ErrorIs(t, driver.Create(ctx, brand, views.Brand), domain.ErrAlreadyExists)
		}
		_ = tx1.Rollback()

		assert.ErrorIs(
			t,
			driver.Update(
				ctx,
				&domain.Brand{Id: "not_exist", Title: "Test Brand"},
				views.Brand,
			),
			domain.ErrNotFound,
		)

		assert.ErrorIs(t, driver.Delete(ctx, brand.Id, views.Brand), domain.ErrNotFound)
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

		assert.ErrorIs(t, driver.Create(ctx, 2, views.Category), domain.ErrInvalidType)
		assert.ErrorIs(t, driver.Create(ctx, 2, 19), domain.ErrUnknownType)

		cat := &domain.Category{Id: "cat_test_bad", Title: "Test Cat", Uri: "test-uri"}

		tx1, err := db.Driver.Begin()
		assert.NoError(t, err)
		{
			driver := Driver{Driver: tx1}
			assert.NoError(t, driver.Create(ctx, cat, views.Category))
			assert.ErrorIs(t, driver.Create(ctx, cat, views.Category), domain.ErrAlreadyExists)
		}
		_ = tx1.Rollback()

		assert.ErrorIs(
			t,
			driver.Update(
				ctx,
				&domain.Category{Id: "not_exist", Title: "Test Cat", Uri: "test-uri"},
				views.Category,
			),
			domain.ErrNotFound,
		)

		assert.ErrorIs(t, driver.Delete(ctx, cat.Id, views.Category), domain.ErrNotFound)
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

		assert.ErrorIs(t, driver.Create(ctx, 2, views.Country), domain.ErrInvalidType)
		assert.ErrorIs(t, driver.Create(ctx, 2, 19), domain.ErrUnknownType)

		country := &domain.Country{Id: "country_test_bad", Title: "Test Country", Friendly: "FriendlyName"}

		tx1, err := db.Driver.Begin()
		assert.NoError(t, err)
		{
			driver := Driver{Driver: tx1}
			assert.NoError(t, driver.Create(ctx, country, views.Country))
			assert.ErrorIs(t, driver.Create(ctx, country, views.Country), domain.ErrAlreadyExists)
		}
		_ = tx1.Rollback()

		assert.ErrorIs(
			t,
			driver.Update(
				ctx,
				&domain.Country{Id: "not_exist", Title: "Test Country", Friendly: "FriendlyName"},
				views.Country,
			),
			domain.ErrNotFound,
		)

		assert.ErrorIs(t, driver.Delete(ctx, country.Id, views.Country), domain.ErrNotFound)
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

		assert.ErrorIs(t, driver.Create(ctx, 2, views.Material), domain.ErrInvalidType)
		assert.ErrorIs(t, driver.Create(ctx, 2, 19), domain.ErrUnknownType)

		mat := &domain.Material{Id: "mat_test_bad", Title: "Test Material"}
		tx1, err := db.Driver.Begin()
		assert.NoError(t, err)
		{
			driver := Driver{Driver: tx1}
			assert.NoError(t, driver.Create(ctx, mat, views.Material))
			assert.ErrorIs(t, driver.Create(ctx, mat, views.Material), domain.ErrAlreadyExists)
		}
		_ = tx1.Rollback()

		assert.ErrorIs(
			t,
			driver.Update(
				ctx,
				&domain.Material{Id: "not_exist", Title: "Test Material"},
				views.Material,
			),
			domain.ErrNotFound,
		)

		assert.ErrorIs(t, driver.Delete(ctx, mat.Id, views.Material), domain.ErrNotFound)
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

		assert.ErrorIs(t, driver.Create(ctx, 2, views.Color), domain.ErrInvalidType)
		assert.ErrorIs(t, driver.Create(ctx, 2, 19), domain.ErrUnknownType)

		color := &domain.Color{Id: "color_test_bad", Title: "TestColor", Hex: "#123456"}
		tx1, err := db.Driver.Begin()
		assert.NoError(t, err)
		{
			driver := Driver{Driver: tx1}
			assert.NoError(t, driver.Create(ctx, color, views.Color))
			assert.ErrorIs(t, driver.Create(ctx, color, views.Color), domain.ErrAlreadyExists)
		}
		_ = tx1.Rollback()

		assert.ErrorIs(
			t,
			driver.Update(
				ctx,
				&domain.Color{Id: "not_exist", Title: "TestColor", Hex: "#123456"},
				views.Color,
			),
			domain.ErrNotFound,
		)

		assert.ErrorIs(t, driver.Delete(ctx, color.Id, views.Color), domain.ErrNotFound)
	})

	t.Run("slides bad", func(t *testing.T) {
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

		assert.ErrorIs(t, driver.Create(ctx, 2, views.Slide), domain.ErrInvalidType)
		assert.ErrorIs(t, driver.Create(ctx, 2, 19), domain.ErrUnknownType)

		s := &domain.Slide{Id: "slide test", Link: "http://example.com", Img: "img1.jpg", Img762: "img762.jpg"}
		tx1, err := db.Driver.Begin()
		assert.NoError(t, err)
		{
			driver := Driver{Driver: tx1}
			assert.NoError(t, driver.Create(ctx, s, views.Slide))
			assert.ErrorIs(t, driver.Create(ctx, s, views.Slide), domain.ErrAlreadyExists)
		}
		_ = tx1.Rollback()

		assert.ErrorIs(
			t,
			driver.Update(
				ctx,
				&domain.Slide{Id: "notexist"},
				views.Slide,
			),
			domain.ErrNotFound,
		)

		assert.ErrorIs(t, driver.Delete(ctx, s.Id, views.Slide), domain.ErrNotFound)
	})

	t.Run("articles bad", func(t *testing.T) {
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

		assert.ErrorIs(t, driver.Create(ctx, 2, views.Article), domain.ErrInvalidType)
		assert.ErrorIs(t, driver.Create(ctx, 2, 19), domain.ErrUnknownType)

		s := &domain.Article{Id: "article test", Title: "Title", Img: "img.jpg", Text: "# zagolovok"}
		tx1, err := db.Driver.Begin()
		assert.NoError(t, err)
		{
			driver := Driver{Driver: tx1}
			assert.NoError(t, driver.Create(ctx, s, views.Article))
			assert.ErrorIs(t, driver.Create(ctx, s, views.Article), domain.ErrAlreadyExists)
		}
		_ = tx1.Rollback()

		assert.ErrorIs(
			t,
			driver.Update(
				ctx,
				&domain.Article{Id: "notexist"},
				views.Article,
			),
			domain.ErrNotFound,
		)

		assert.ErrorIs(t, driver.Delete(ctx, s.Id, views.Article), domain.ErrNotFound)
	})
}
