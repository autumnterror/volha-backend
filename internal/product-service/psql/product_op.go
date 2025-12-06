package psql

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/autumnterror/breezynotes/pkg/log"
	"github.com/autumnterror/breezynotes/pkg/utils/format"
	productsRPC "github.com/autumnterror/volha-backend/pkg/proto/gen"
	"github.com/lib/pq"
)

func (d Driver) GetAllProducts(ctx context.Context, start, end int) ([]*productsRPC.Product, error) {
	const op = "PostgresDb.GetAllProducts"

	if end <= start {
		return []*productsRPC.Product{}, nil
	}

	limit := end - start

	rows, err := d.Driver.QueryContext(ctx, getAllProductsQuery, limit, start)
	if err != nil {
		return nil, format.Error(op, err)
	}
	defer rows.Close()

	var products []*productsRPC.Product

	for rows.Next() {
		var (
			materialsJSON, colorsJSON []byte
			similarJSON               []byte
		)
		p := emptyProduct
		if err := rows.Scan(
			&p.Id,
			&p.Title,
			&p.Article,
			&p.Width,
			&p.Height,
			&p.Depth,
			pq.Array(&p.Photos),
			&p.Price,
			&p.Description,

			&p.Brand.Id,
			&p.Brand.Title,

			&p.Category.Id,
			&p.Category.Title,
			&p.Category.Uri,

			&p.Country.Id,
			&p.Country.Title,
			&p.Country.Friendly,

			&materialsJSON,
			&colorsJSON,
			&similarJSON,
		); err != nil {
			log.Error(op, "", err)
			continue
		}

		if err := json.Unmarshal(materialsJSON, &p.Materials); err != nil {
			log.Error(op, "", err)
			continue
		}
		if err := json.Unmarshal(colorsJSON, &p.Colors); err != nil {
			log.Error(op, "", err)
			continue
		}
		if err := json.Unmarshal(similarJSON, &p.Seems); err != nil {
			log.Error(op, "", err)
			continue
		}

		products = append(products, &p)
	}

	if err := rows.Err(); err != nil {
		return nil, format.Error(op, err)
	}

	return products, nil
}

func (d Driver) GetProduct(ctx context.Context, id string) (*productsRPC.Product, error) {
	const op = "PostgresDb.GetProduct"

	var (
		materialsJSON, colorsJSON []byte
		similarJSON               []byte
	)
	p := emptyProduct
	err := d.Driver.QueryRowContext(ctx, getProductQuery, id).Scan(
		&p.Id, &p.Title, &p.Article, &p.Width, &p.Height, &p.Depth,
		pq.Array(&p.Photos), &p.Price, &p.Description,
		&p.Brand.Id, &p.Brand.Title,
		&p.Category.Id, &p.Category.Title, &p.Category.Uri,
		&p.Country.Id, &p.Country.Title, &p.Country.Friendly,
		&materialsJSON, &colorsJSON, &similarJSON,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, format.Error(op, err)
	}

	if err := json.Unmarshal(materialsJSON, &p.Materials); err != nil {
		return nil, format.Error(op, err)
	}

	if err := json.Unmarshal(colorsJSON, &p.Colors); err != nil {
		return nil, format.Error(op, err)
	}

	if err := json.Unmarshal(similarJSON, &p.Seems); err != nil {
		return nil, format.Error(op, err)
	}

	return &p, nil
}

func (d Driver) CreateProduct(ctx context.Context, p *productsRPC.ProductId) error {
	const op = "PostgresDb.CreateProduct"

	_, err := d.Driver.ExecContext(
		ctx,
		createProductQuery,
		p.Id, p.Title, p.Article, p.Brand, p.Category, p.Country, p.Width, p.Height, p.Depth,
		pq.Array(p.Photos), p.Price, p.Description,
		pq.Array(p.Materials), pq.Array(p.Colors), pq.Array(p.Seems),
	)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			switch pqErr.Code {
			case "23503": // foreign_key_violation
				log.Error(op, pqErr.Message, nil)
				return format.Error(op, ErrForeignKey)
			case "23505": // unique_violation
				log.Error(op, pqErr.Message, nil)
				return format.Error(op, ErrAlreadyExists)
			}
		}
		return format.Error(op, err)
	}

	return nil
}

func (d Driver) UpdateProduct(ctx context.Context, p *productsRPC.ProductId) error {
	const op = "PostgresDb.UpdateProduct"

	res, err := d.Driver.ExecContext(ctx, updateProductQuery,
		p.Id, p.Title, p.Article, p.Brand, p.Category, p.Country, p.Width, p.Height, p.Depth,
		pq.Array(p.Photos), p.Price, p.Description,
		pq.Array(p.Materials), pq.Array(p.Colors), pq.Array(p.Seems),
	)
	if err != nil {
		return format.Error(op, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return format.Error(op, err)
	}
	if rowsAffected == 0 {
		return format.Error(op, ErrNotFound)
	}

	return nil
}

func (d Driver) DeleteProduct(ctx context.Context, id string) error {
	const op = "PostgresDb.DeleteProduct"
	res, err := d.Driver.ExecContext(ctx, `DELETE FROM products WHERE id = $1`, id)
	if err != nil {
		return format.Error(op, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return format.Error(op, err)
	}
	if rowsAffected == 0 {
		return format.Error(op, ErrNotFound)
	}
	return nil
}

func (d Driver) SearchProducts(ctx context.Context, filter *productsRPC.ProductSearch) ([]*productsRPC.Product, error) {
	const op = "PostgresDb.SearchProducts"

	var (
		query string
		args  []any
	)
	query = filterProductsQuery
	switch {
	case filter.Id != "":
		query += ` WHERE id = $1`
		args = append(args, filter.Id)
	case filter.Article != "":
		query += ` WHERE article = $1`
		args = append(args, filter.Article)
	case filter.Title != "":
		query += ` WHERE p.title ILIKE $1`
		args = append(args, "%"+filter.Title+"%")
	default:
		return nil, format.Error(op, errors.New("no search parameter provided"))
	}

	rows, err := d.Driver.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, format.Error(op, err)
	}
	defer rows.Close()

	var products []*productsRPC.Product

	for rows.Next() {
		var (
			materialsJSON, colorsJSON []byte
			similarJSON               []byte
		)
		p := emptyProduct
		if err := rows.Scan(
			&p.Id, &p.Title, &p.Article, &p.Width, &p.Height, &p.Depth,
			pq.Array(&p.Photos), &p.Price, &p.Description,
			&p.Brand.Id, &p.Brand.Title,
			&p.Category.Id, &p.Category.Title, &p.Category.Uri,
			&p.Country.Id, &p.Country.Title, &p.Country.Friendly,
			&materialsJSON, &colorsJSON, &similarJSON,
		); err != nil {
			log.Error(op, "", err)
			continue
		}

		if err := json.Unmarshal(materialsJSON, &p.Materials); err != nil {
			log.Error(op, "", err)
			continue
		}
		if err := json.Unmarshal(colorsJSON, &p.Colors); err != nil {
			log.Error(op, "", err)
			continue
		}
		if err := json.Unmarshal(similarJSON, &p.Seems); err != nil {
			log.Error(op, "", err)
			continue
		}

		products = append(products, &p)
	}

	if err := rows.Err(); err != nil {
		return nil, format.Error(op, err)
	}

	return products, nil
}
