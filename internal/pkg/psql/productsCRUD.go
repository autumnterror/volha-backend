package psql

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"log"
	"productService/internal/utils/format"
	"productService/internal/views"
)

type SqlRepo interface {
	Query(query string, args ...any) (*sql.Rows, error)
	Exec(query string, args ...any) (sql.Result, error)
	QueryRow(query string, args ...any) *sql.Row
}

type ProductRepo interface {
	GetAll() ([]views.Product, error)
	Create(p *views.Product) error
	Update(p *views.Product, id string) error
	Delete(id string) error
	FilterProducts(filter *views.ProductFilter) ([]views.Product, error)
}

type Driver struct {
	Driver SqlRepo
}

// GetAll products
func (d Driver) GetAll() ([]views.Product, error) {
	const op = "PostgresDb.GetAll"

	var lp []views.Product
	rows, err := d.Driver.Query(`SELECT * FROM products`)
	if err != nil {
		return nil, format.Error(op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var p views.Product
		if err := rows.Scan(
			&p.Id,
			&p.Title,
			&p.Article,
			&p.Brand,
			&p.Country,
			&p.Width,
			&p.Height,
			&p.Depth,
			pq.Array(&p.Materials),
			pq.Array(&p.Colors),
			pq.Array(&p.Photos),
			pq.Array(&p.Seems),
			&p.Price,
			&p.Description,
		); err != nil {
			log.Println(format.Error(op, err))
			continue
		}
		lp = append(lp, p)
	}

	return lp, nil
}

// Create product
func (d Driver) Create(p *views.Product) error {
	const op = "PostgresDb.Create"

	query := `
        INSERT INTO products (
            id, title, article, brand, country, width, height, depth, 
            materials, color, photos, seems, price, description
        ) 
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`

	_, err := d.Driver.Exec(
		query,
		p.Id,
		&p.Title,
		&p.Article,
		p.Brand,
		p.Country,
		p.Width,
		p.Height,
		p.Depth,
		pq.Array(p.Materials),
		pq.Array(p.Colors),
		pq.Array(p.Photos),
		pq.Array(p.Seems),
		p.Price,
		p.Description,
	)
	if err != nil {
		return format.Error(op, err)
	}

	return nil
}

// Update product
func (d Driver) Update(p *views.Product, id string) error {
	const op = "PostgresDb.Update"

	query := `
        UPDATE products SET
			title = $2,
			article = $3,
            brand = $4,
            country = $5,
            width = $6,
            height = $7,
            depth = $8,
            materials = $9,
            color = $10,
            photos = $11,
            seems = $12,
            price = $13,
            description = $14
        WHERE id = $1
    `

	result, err := d.Driver.Exec(
		query,
		p.Id,
		p.Title,
		p.Article,
		p.Brand,
		p.Country,
		p.Width,
		p.Height,
		p.Depth,
		pq.Array(p.Materials),
		pq.Array(p.Colors),
		pq.Array(p.Photos),
		pq.Array(p.Seems),
		p.Price,
		p.Description,
	)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return format.Error(op, err)
	}

	if rowsAffected == 0 {
		return format.Error(op, fmt.Errorf("p with id %s not found", p.Id))
	}
	return nil
}

// Delete product
func (d Driver) Delete(id string) error {
	const op = "PostgresDb.Delete"

	query := `
				DELETE FROM products
				WHERE id = $1
			`
	if _, err := d.Driver.Exec(query, id); err != nil {
		return format.Error(op, err)
	}

	return nil
}
