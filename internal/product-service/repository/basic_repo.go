package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/autumnterror/volha-backend/pkg/views"

	"github.com/autumnterror/volha-backend/internal/product-service/domain"

	"github.com/autumnterror/breezynotes/pkg/log"
	"github.com/lib/pq"

	"github.com/autumnterror/breezynotes/pkg/utils/format"
)

type BasicRepo interface {
	GetAll(ctx context.Context, _type views.Type) (any, error)
	Get(ctx context.Context, id string, _type views.Type) (any, error)
	Create(ctx context.Context, obj any, _type views.Type) error
	Update(ctx context.Context, obj any, _type views.Type) error
	Delete(ctx context.Context, id string, _type views.Type) error
}

func (d Driver) GetAll(ctx context.Context, _type views.Type) (any, error) {
	op := "psql.GetAll." + _type.String()

	var scanner entityScanner
	var query string

	switch _type {
	case views.Brand:
		scanner = &brandScanner{}
		query = `SELECT id, title FROM brands`
	case views.Country:
		scanner = &countryScanner{}
		query = `SELECT id, title, friendly FROM countries`
	case views.Material:
		scanner = &materialScanner{}
		query = `SELECT id, title FROM materials`
	case views.Color:
		scanner = &colorScanner{}
		query = `SELECT id, title, hex FROM colors`
	case views.Category:
		scanner = &categoryScanner{}
		query = `SELECT id, title, uri, img FROM categories`
	case views.Slide:
		scanner = &slideScanner{}
		query = `SELECT id, link, img, img762 FROM slides`
	case views.Article:
		scanner = &articleScanner{}
		query = `SELECT id, title, img, text FROM articles`
	default:
		return nil, format.Error(op, domain.ErrUnknownType)
	}

	rows, err := d.Driver.QueryContext(ctx, query)
	if err != nil {
		return nil, format.Error(op, err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := scanner.Scan(rows); err != nil {
			log.Error(op, "", err)
			continue
		}
	}

	return scanner.GetList(), nil
}

func (d Driver) Get(
	ctx context.Context,
	id string,
	_type views.Type,
) (any, error) {
	op := "psql.Get." + _type.String()
	var scanner entityScannerRow
	var query string

	switch _type {
	case views.Brand:
		scanner = &brandScannerRow{}
		query = `SELECT id, title FROM brands WHERE id = $1`
	case views.Country:
		scanner = &countryScannerRow{}
		query = `SELECT id, title, friendly FROM countries WHERE id = $1`
	case views.Material:
		scanner = &materialScannerRow{}
		query = `SELECT id, title FROM materials WHERE id = $1`
	case views.Color:
		scanner = &colorScannerRow{}
		query = `SELECT id, title, hex FROM colors WHERE id = $1`
	case views.Category:
		scanner = &categoryScannerRow{}
		query = `SELECT id, title, uri, img FROM categories WHERE id = $1`
	case views.Slide:
		scanner = &slideScannerRow{}
		query = `SELECT id, link, img, img762 FROM slides WHERE id = $1`
	case views.Article:
		scanner = &articleScannerRow{}
		query = `SELECT id, title, img, text FROM articles WHERE id = $1`
	default:
		return nil, format.Error(op, domain.ErrUnknownType)
	}
	res := d.Driver.QueryRowContext(ctx, query, id)
	if err := scanner.Scan(res); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, format.Error(op, domain.ErrNotFound)
		}
		return nil, format.Error(op, err)
	}
	return scanner.Get(), nil
}

func (d Driver) Create(
	ctx context.Context,
	obj any,
	_type views.Type,
) error {
	op := "psql.Create." + _type.String()
	var query string
	var args []any
	switch _type {
	case views.Brand:
		b, ok := obj.(*domain.Brand)
		if !ok {
			return format.Error(op, domain.ErrInvalidType)
		}
		query = `INSERT INTO brands (id, title) VALUES ($1, $2)`
		args = append(args, b.Id)
		args = append(args, b.Title)
	case views.Country:
		b, ok := obj.(*domain.Country)
		if !ok {
			return format.Error(op, domain.ErrInvalidType)
		}
		query = `INSERT INTO countries (id, title, friendly) VALUES ($1, $2, $3)`
		args = append(args, b.Id)
		args = append(args, b.Title)
		args = append(args, b.Friendly)
	case views.Material:
		b, ok := obj.(*domain.Material)
		if !ok {
			return format.Error(op, domain.ErrInvalidType)
		}
		query = `INSERT INTO materials (id, title) VALUES ($1, $2)`
		args = append(args, b.Id)
		args = append(args, b.Title)
	case views.Category:
		b, ok := obj.(*domain.Category)
		if !ok {
			return format.Error(op, domain.ErrInvalidType)
		}
		query = `INSERT INTO categories (id, title, uri, img) VALUES ($1, $2, $3, $4)`
		args = append(args, b.Id)
		args = append(args, b.Title)
		args = append(args, b.Uri)
		args = append(args, b.Img)
	case views.Color:
		b, ok := obj.(*domain.Color)
		if !ok {
			return format.Error(op, domain.ErrInvalidType)
		}
		query = `INSERT INTO colors (id, title, hex) VALUES ($1, $2, $3)`
		args = append(args, b.Id)
		args = append(args, b.Title)
		args = append(args, b.Hex)
	case views.Slide:
		b, ok := obj.(*domain.Slide)
		if !ok {
			return format.Error(op, domain.ErrInvalidType)
		}
		query = `INSERT INTO slides (id, link, img, img762) VALUES ($1, $2, $3, $4)`
		args = append(args, b.Id)
		args = append(args, b.Link)
		args = append(args, b.Img)
		args = append(args, b.Img762)
	case views.Article:
		b, ok := obj.(*domain.Article)
		if !ok {
			return format.Error(op, domain.ErrInvalidType)
		}
		query = `INSERT INTO articles (id, title, img, text) VALUES ($1, $2, $3, $4)`
		args = append(args, b.Id)
		args = append(args, b.Title)
		args = append(args, b.Img)
		args = append(args, b.Text)
	default:
		return format.Error(op, domain.ErrUnknownType)
	}

	_, err := d.Driver.ExecContext(ctx, query, args...)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			switch pqErr.Code {
			case "23505": // unique_violation
				return format.Error(op, domain.ErrAlreadyExists)
			case "23503": // foreign_key_violation
				return format.Error(op, domain.ErrForeignKey)
			}
		}
		return format.Error(op, err)
	}
	return nil
}

// Update can return domain.ErrNotFound
func (d Driver) Update(
	ctx context.Context,
	obj any,
	_type views.Type,
) error {
	op := "psql.Update." + _type.String()

	var query string
	var args []any

	switch _type {
	case views.Brand:
		b, ok := obj.(*domain.Brand)
		if !ok {
			return format.Error(op, domain.ErrInvalidType)
		}
		query = `UPDATE brands SET title = $2 WHERE id = $1`
		args = append(args, b.Id)
		args = append(args, b.Title)
	case views.Country:
		b, ok := obj.(*domain.Country)
		if !ok {
			return format.Error(op, domain.ErrInvalidType)
		}
		query = `UPDATE countries SET title = $2, friendly = $3 WHERE id = $1`
		args = append(args, b.Id)
		args = append(args, b.Title)
		args = append(args, b.Friendly)
	case views.Material:
		b, ok := obj.(*domain.Material)
		if !ok {
			return format.Error(op, domain.ErrInvalidType)
		}
		query = `UPDATE materials SET title = $2 WHERE id = $1`
		args = append(args, b.Id)
		args = append(args, b.Title)
	case views.Category:
		b, ok := obj.(*domain.Category)
		if !ok {
			return format.Error(op, domain.ErrInvalidType)
		}
		query = `UPDATE categories SET title = $2, uri = $3, img = $4 WHERE id = $1`
		args = append(args, b.Id)
		args = append(args, b.Title)
		args = append(args, b.Uri)
		args = append(args, b.Img)
	case views.Color:
		b, ok := obj.(*domain.Color)
		if !ok {
			return format.Error(op, domain.ErrInvalidType)
		}
		query = `UPDATE colors SET title = $2, hex = $3 WHERE id = $1`
		args = append(args, b.Id)
		args = append(args, b.Title)
		args = append(args, b.Hex)
	case views.Slide:
		b, ok := obj.(*domain.Slide)
		if !ok {
			return format.Error(op, domain.ErrInvalidType)
		}
		query = `UPDATE slides SET link = $2, img = $3, img762 = $4 WHERE id = $1`
		args = append(args, b.Id)
		args = append(args, b.Link)
		args = append(args, b.Img)
		args = append(args, b.Img762)
	case views.Article:
		b, ok := obj.(*domain.Article)
		if !ok {
			return format.Error(op, domain.ErrInvalidType)
		}
		query = `UPDATE articles SET title = $2, img = $3, text = $4 WHERE id = $1`
		args = append(args, b.Id)
		args = append(args, b.Title)
		args = append(args, b.Img)
		args = append(args, b.Text)
	default:
		return format.Error(op, domain.ErrUnknownType)
	}

	res, err := d.Driver.ExecContext(ctx, query, args...)
	if err != nil {
		return format.Error(op, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return format.Error(op, err)
	}
	if rowsAffected == 0 {
		return format.Error(op, domain.ErrNotFound)
	}

	return nil
}

func (d Driver) Delete(
	ctx context.Context,
	id string,
	_type views.Type,
) error {
	op := "psql.Delete." + _type.String()
	var query string
	switch _type {
	case views.Brand:
		query = `DELETE FROM brands WHERE id = $1`
	case views.Country:
		query = `DELETE FROM countries WHERE id = $1`
	case views.Material:
		query = `DELETE FROM materials WHERE id = $1`
	case views.Category:
		query = `DELETE FROM categories WHERE id = $1`
	case views.Color:
		query = `DELETE FROM colors WHERE id = $1`
	case views.Slide:
		query = `DELETE FROM slides WHERE id = $1`
	case views.Article:
		query = `DELETE FROM articles WHERE id = $1`
	default:
		return format.Error(op, domain.ErrUnknownType)
	}

	res, err := d.Driver.ExecContext(ctx, query, id)
	if err != nil {
		return format.Error(op, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return format.Error(op, err)
	}
	if rowsAffected == 0 {
		return format.Error(op, domain.ErrNotFound)
	}
	return nil
}
