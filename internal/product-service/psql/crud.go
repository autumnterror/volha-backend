package psql

import (
	"context"
	"log"

	"github.com/autumnterror/breezynotes/pkg/utils/format"
	productsRPC "github.com/autumnterror/volha-backend/pkg/proto/gen"
	"github.com/autumnterror/volha-backend/pkg/views"
)

func (d Driver) GetAll(ctx context.Context, _type views.Type) (any, error) {
	op := "psql.GetAll." + _type.String()

	var scanner EntityScanner
	var query string

	switch _type {
	case views.Brand:
		scanner = &BrandScanner{}
		query = `SELECT id, title FROM brands`
	case views.Country:
		scanner = &CountryScanner{}
		query = `SELECT id, title, friendly FROM countries`
	case views.Material:
		scanner = &MaterialScanner{}
		query = `SELECT id, title FROM materials`
	case views.Color:
		scanner = &ColorScanner{}
		query = `SELECT id, title, hex FROM colors`
	case views.Category:
		scanner = &CategoryScanner{}
		query = `SELECT id, title, uri, img FROM categories`

	default:
		return nil, format.Error(op, ErrUnknownType)
	}

	rows, err := d.Driver.QueryContext(ctx, query)
	if err != nil {
		return nil, format.Error(op, err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := scanner.Scan(rows); err != nil {
			log.Println(format.Error(op, err))
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
	var scanner EntityScannerRow
	var query string

	switch _type {
	case views.Brand:
		scanner = &BrandScannerRow{}
		query = `SELECT id, title FROM brands WHERE id = $1`
	case views.Country:
		scanner = &CountryScannerRow{}
		query = `SELECT id, title, friendly FROM countries WHERE id = $1`
	case views.Material:
		scanner = &MaterialScannerRow{}
		query = `SELECT id, title FROM materials WHERE id = $1`
	case views.Color:
		scanner = &ColorScannerRow{}
		query = `SELECT id, title, hex FROM colors WHERE id = $1`
	case views.Category:
		scanner = &CategoryScannerRow{}
		query = `SELECT id, title, uri, img FROM categories WHERE id = $1`
	default:
		return nil, format.Error(op, ErrUnknownType)
	}
	res := d.Driver.QueryRowContext(ctx, query, id)
	if err := scanner.Scan(res); err != nil {
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
	args := []any{}
	switch _type {
	case views.Brand:
		b, ok := obj.(*productsRPC.Brand)
		if !ok {
			return format.Error(op, ErrInvalidType)
		}
		query = `INSERT INTO brands (id, title) VALUES ($1, $2)`
		args = append(args, b.Id)
		args = append(args, b.Title)
	case views.Country:
		b, ok := obj.(*productsRPC.Country)
		if !ok {
			return format.Error(op, ErrInvalidType)
		}
		query = `INSERT INTO countries (id, title, friendly) VALUES ($1, $2, $3)`
		args = append(args, b.Id)
		args = append(args, b.Title)
		args = append(args, b.Friendly)
	case views.Material:
		b, ok := obj.(*productsRPC.Material)
		if !ok {
			return format.Error(op, ErrInvalidType)
		}
		query = `INSERT INTO materials (id, title) VALUES ($1, $2)`
		args = append(args, b.Id)
		args = append(args, b.Title)
	case views.Category:
		b, ok := obj.(*productsRPC.Category)
		if !ok {
			return format.Error(op, ErrInvalidType)
		}
		query = `INSERT INTO categories (id, title, uri, img) VALUES ($1, $2, $3, $4)`
		args = append(args, b.Id)
		args = append(args, b.Title)
		args = append(args, b.Uri)
		args = append(args, b.Img)
	case views.Color:
		b, ok := obj.(*productsRPC.Color)
		if !ok {
			return format.Error(op, ErrInvalidType)
		}
		query = `INSERT INTO colors (id, title, hex) VALUES ($1, $2, $3)`
		args = append(args, b.Id)
		args = append(args, b.Title)
		args = append(args, b.Hex)
	default:
		return format.Error(op, ErrUnknownType)
	}

	_, err := d.Driver.ExecContext(ctx, query, args...)
	if err != nil {
		return format.Error(op, err)
	}
	return nil
}

func (d Driver) Update(
	ctx context.Context,
	id string,
	obj any,
	_type views.Type,
) error {
	op := "psql.Update." + _type.String()
	switch _type {
	case views.Brand:
		b, ok := obj.(*productsRPC.Brand)
		if !ok {
			return format.Error(op, ErrInvalidType)
		}
		query := `UPDATE brands SET title = $2 WHERE id = $1`
		result, err := d.Driver.ExecContext(ctx, query, id, b.title)
		if err != nil {
			return format.Error(op, err)
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return format.Error(op, err)
		}
		if rowsAffected == 0 {
			return format.Error(op, ErrNotFound)
		}

		return nil
	case views.Country:
		return nil
	default:
		return format.Error(op, ErrUnknownType)
	}
}

func (d Driver) Delete(
	ctx context.Context,
	id string,
	_type views.Type,
) error {
	op := "psql.Delete." + _type.String()
	switch _type {
	case views.Brand:
		query := `DELETE FROM brands WHERE id = $1`
		_, err := d.Driver.ExecContext(ctx, query, id)
		if err != nil {
			return format.Error(op, err)
		}
		return nil
	case views.Country:
		return nil
	default:
		return format.Error(op, ErrUnknownType)
	}
}
