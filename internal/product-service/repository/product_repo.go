package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/autumnterror/breezynotes/pkg/log"
	"github.com/autumnterror/breezynotes/pkg/utils/format"
	"github.com/autumnterror/volha-backend/internal/product-service/domain"
	"github.com/lib/pq"
)

type ProductRepo interface {
	GetAllProducts(ctx context.Context, start, end int) ([]*domain.Product, error)
	GetProduct(ctx context.Context, id string) (*domain.Product, error)
	CreateProduct(ctx context.Context, p *domain.ProductId) error
	UpdateProduct(ctx context.Context, p *domain.ProductId) error
	IncrementViews(ctx context.Context, id string) error
	DeleteProduct(ctx context.Context, id string) error
	SearchProducts(ctx context.Context, filter *domain.ProductSearch) ([]*domain.Product, error)
	FilterProducts(ctx context.Context, filter *domain.ProductFilter) ([]*domain.Product, error)
}

func (d Driver) GetAllProducts(ctx context.Context, start, end int) ([]*domain.Product, error) {
	const op = "PostgresDb.GetAllProducts"

	if end <= start {
		return []*domain.Product{}, nil
	}

	limit := end - start

	rows, err := d.Driver.QueryContext(ctx, getAllProductsQuery, limit, start)
	if err != nil {
		return nil, format.Error(op, err)
	}
	defer rows.Close()

	var products []*domain.Product

	for rows.Next() {
		var (
			materialsJSON, colorsJSON []byte
			similarJSON               []byte
		)
		p := domain.NewEmptyProduct()
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
			&p.Views,

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

func (d Driver) GetProduct(ctx context.Context, id string) (*domain.Product, error) {
	const op = "PostgresDb.GetProduct"

	var (
		materialsJSON, colorsJSON []byte
		similarJSON               []byte
	)
	p := domain.NewEmptyProduct()
	err := d.Driver.QueryRowContext(ctx, getProductQuery, id).Scan(
		&p.Id, &p.Title, &p.Article, &p.Width, &p.Height, &p.Depth,
		pq.Array(&p.Photos), &p.Price, &p.Description, &p.Views,
		&p.Brand.Id, &p.Brand.Title,
		&p.Category.Id, &p.Category.Title, &p.Category.Uri,
		&p.Country.Id, &p.Country.Title, &p.Country.Friendly,
		&materialsJSON, &colorsJSON, &similarJSON,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
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

func (d Driver) CreateProduct(ctx context.Context, p *domain.ProductId) error {
	const op = "PostgresDb.CreateProduct"

	_, err := d.Driver.ExecContext(
		ctx,
		createProductQuery,
		p.Id, p.Title, p.Article, p.Brand, p.Category, p.Country, p.Width, p.Height, p.Depth,
		pq.Array(p.Photos), p.Price, p.Description, p.Views,
		pq.Array(p.Materials), pq.Array(p.Colors), pq.Array(p.Seems),
	)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			switch pqErr.Code {
			case "23503": // foreign_key_violation
				log.Error(op, pqErr.Message, nil)
				return format.Error(op, domain.ErrForeignKey)
			case "23505": // unique_violation
				log.Error(op, pqErr.Message, nil)
				return format.Error(op, domain.ErrAlreadyExists)
			}
		}
		return format.Error(op, err)
	}

	return nil
}

func (d Driver) UpdateProduct(ctx context.Context, p *domain.ProductId) error {
	const op = "PostgresDb.UpdateProduct"

	res, err := d.Driver.ExecContext(ctx, updateProductQuery,
		p.Id, p.Title, p.Article, p.Brand, p.Category, p.Country, p.Width, p.Height, p.Depth,
		pq.Array(p.Photos), p.Price, p.Description, p.Views,
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
		return format.Error(op, domain.ErrNotFound)
	}

	return nil
}

func (d Driver) IncrementViews(ctx context.Context, id string) error {
	const op = "PostgresDb.IncrementViews"

	query := `UPDATE products SET views = views + 1 WHERE id = $1`

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
		return format.Error(op, domain.ErrNotFound)
	}
	return nil
}

func (d Driver) SearchProducts(ctx context.Context, filter *domain.ProductSearch) ([]*domain.Product, error) {
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

	var products []*domain.Product

	for rows.Next() {
		var (
			materialsJSON, colorsJSON []byte
			similarJSON               []byte
		)
		p := domain.NewEmptyProduct()
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

func (d Driver) FilterProducts(ctx context.Context, filter *domain.ProductFilter) ([]*domain.Product, error) {
	const op = "PostgresDb.FilterProducts"

	var conditions []string
	var args []any
	argPos := 1

	buildFilter := func(field string, values []string) {
		if len(values) == 0 {
			return
		}
		placeholders := make([]string, len(values))
		for i, v := range values {
			placeholders[i] = fmt.Sprintf("$%d", argPos)
			args = append(args, v)
			argPos++
		}
		conditions = append(conditions, fmt.Sprintf("%s IN (%s)", field, strings.Join(placeholders, ",")))
	}

	buildFilter("brand_id", filter.Brand)
	buildFilter("category_id", filter.Category)
	buildFilter("country_id", filter.Country)

	numericFilters := []struct {
		field string
		value int
		op    string
	}{
		{"width", int(filter.MinWidth), ">="},
		{"width", int(filter.MaxWidth), "<="},
		{"height", int(filter.MinHeight), ">="},
		{"height", int(filter.MaxHeight), "<="},
		{"depth", int(filter.MinDepth), ">="},
		{"depth", int(filter.MaxDepth), "<="},
		{"price", int(filter.MinPrice), ">="},
		{"price", int(filter.MaxPrice), "<="},
	}

	for _, f := range numericFilters {
		if f.value > 0 {
			conditions = append(conditions, fmt.Sprintf("%s %s $%d", f.field, f.op, argPos))
			args = append(args, f.value)
			argPos++
		}
	}

	if len(filter.Materials) > 0 {
		placeholders := make([]string, len(filter.Materials))
		for i, v := range filter.Materials {
			placeholders[i] = fmt.Sprintf("$%d", argPos)
			args = append(args, v)
			argPos++
		}
		conditions = append(conditions, fmt.Sprintf(`
        EXISTS (
            SELECT 1 FROM product_materials pm
            WHERE pm.product_id = p.id AND pm.material_id IN (%s)
        )`, strings.Join(placeholders, ",")))
	}

	if len(filter.Colors) > 0 {
		placeholders := make([]string, len(filter.Colors))
		for i, v := range filter.Colors {
			placeholders[i] = fmt.Sprintf("$%d", argPos)
			args = append(args, v)
			argPos++
		}
		conditions = append(conditions, fmt.Sprintf(`
        EXISTS (
            SELECT 1 FROM product_colors pc
            WHERE pc.product_id = p.id AND pc.color_id IN (%s)
        )`, strings.Join(placeholders, ",")))
	}

	query := filterProductsQuery
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	if filter.SortBy != "" {
		validSort := map[string]bool{"price": true, "width": true, "height": true, "depth": true, "title": true, "views": true}
		if validSort[filter.SortBy] {
			order := "ASC"
			if strings.ToLower(filter.SortOrder) == "desc" {
				order = "DESC"
			}
			query += fmt.Sprintf(" ORDER BY %s %s", filter.SortBy, order)
		}
	}

	if filter.Limit > 0 {
		query += fmt.Sprintf(" LIMIT $%d", argPos)
		args = append(args, filter.Limit)
		argPos++
	}
	if filter.Offset > 0 {
		query += fmt.Sprintf(" OFFSET $%d", argPos)
		args = append(args, filter.Offset)
	}

	rows, err := d.Driver.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, format.Error(op, err)
	}
	defer rows.Close()

	var products []*domain.Product

	for rows.Next() {
		var (
			materialsJSON, colorsJSON []byte
			similarJSON               []byte
		)
		p := domain.NewEmptyProduct()
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
			&p.Views,

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
