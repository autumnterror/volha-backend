package psql

import (
	"fmt"
	"github.com/lib/pq"
	"productService/internal/views"
	"strings"
)

func (d Driver) FilterProducts(filter *views.ProductFilter) ([]views.Product, error) {
	//func FilterProducts(filter ProductFilter) ([]views.Product, error) {

	const op = "PostgresDb.FilterProducts"

	var products []views.Product
	var conditions []string
	var args []interface{}
	argPos := 1

	if len(filter.Brand) > 0 {
		var brandConditions []string
		for _, brand := range filter.Brand {
			brandConditions = append(brandConditions, fmt.Sprintf("brand ILIKE $%d", argPos))
			args = append(args, "%"+brand+"%")
			argPos++
		}
		conditions = append(conditions, "("+strings.Join(brandConditions, " OR ")+")")
	}

	if len(filter.Country) > 0 {
		var countryConditions []string
		for _, country := range filter.Country {
			countryConditions = append(countryConditions, fmt.Sprintf("country ILIKE $%d", argPos))
			args = append(args, "%"+country+"%")
			argPos++
		}
		conditions = append(conditions, "("+strings.Join(countryConditions, " OR ")+")")
	}

	if filter.MinWidth > 0 {
		conditions = append(conditions, fmt.Sprintf("width >= $%d", argPos))
		args = append(args, filter.MinWidth)
		argPos++
	}

	if filter.MaxWidth > 0 {
		conditions = append(conditions, fmt.Sprintf("width <= $%d", argPos))
		args = append(args, filter.MaxWidth)
		argPos++
	}

	if filter.MinHeight > 0 {
		conditions = append(conditions, fmt.Sprintf("height >= $%d", argPos))
		args = append(args, filter.MinHeight)
		argPos++
	}

	if filter.MaxHeight > 0 {
		conditions = append(conditions, fmt.Sprintf("height <= $%d", argPos))
		args = append(args, filter.MaxHeight)
		argPos++
	}

	if filter.MinDepth > 0 {
		conditions = append(conditions, fmt.Sprintf("depth >= $%d", argPos))
		args = append(args, filter.MinDepth)
		argPos++
	}

	if filter.MaxDepth > 0 {
		conditions = append(conditions, fmt.Sprintf("depth <= $%d", argPos))
		args = append(args, filter.MaxDepth)
		argPos++
	}

	if len(filter.Materials) > 0 {
		conditions = append(conditions, fmt.Sprintf("materials @> $%d", argPos))
		args = append(args, pq.Array(filter.Materials))
		argPos++
	}

	if len(filter.Colors) > 0 {
		conditions = append(conditions, fmt.Sprintf("color && $%d", argPos))
		args = append(args, pq.Array(filter.Colors))
		argPos++
	}

	if filter.MinPrice > 0 {
		conditions = append(conditions, fmt.Sprintf("price >= $%d", argPos))
		args = append(args, filter.MinPrice)
		argPos++
	}

	if filter.MaxPrice > 0 {
		conditions = append(conditions, fmt.Sprintf("price <= $%d", argPos))
		args = append(args, filter.MaxPrice)
		argPos++
	}

	// Собираем запрос
	query := "SELECT * FROM products"
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	// Добавляем сортировку
	if filter.SortBy != "" {
		validSortFields := map[string]bool{
			"price":  true,
			"width":  true,
			"height": true,
			"depth":  true,
			"title":  true,
		}

		if validSortFields[filter.SortBy] {
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

	rows, err := d.Driver.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var p views.Product
		err := rows.Scan(
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
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return products, nil
}
