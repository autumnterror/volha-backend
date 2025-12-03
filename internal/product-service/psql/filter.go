package psql

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/autumnterror/breezynotes/pkg/log"
	"github.com/autumnterror/breezynotes/pkg/utils/format"
	productsRPC "github.com/autumnterror/volha-backend/pkg/proto/gen"
	"github.com/lib/pq"
	"strings"
)

func (d Driver) FilterProducts(ctx context.Context, filter *productsRPC.ProductFilter) ([]*productsRPC.Product, error) {
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
		validSort := map[string]bool{"price": true, "width": true, "height": true, "depth": true, "title": true}
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
