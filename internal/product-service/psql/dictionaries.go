package psql

import (
	"context"
	"log"
	"strconv"
	"strings"

	"github.com/autumnterror/breezynotes/pkg/utils/format"
	productsRPC "github.com/autumnterror/volha-backend/pkg/proto/gen"
)

func (d Driver) GetDictionariesByCategory(ctx context.Context, id string) (*productsRPC.Dictionaries, error) {
	const op = "PostgresDb.GetDictionaries"

	query := `
		WITH dicts AS (
			SELECT 'brand' as type, id, name, '' as extra1, '' as extra2 FROM brands
			UNION ALL
			SELECT 'category', id, title, uri, '' FROM categories
			UNION ALL
			SELECT 'country', id, title, friendly, '' FROM countries
			UNION ALL
			SELECT 'material', id, title, '', '' FROM materials
			UNION ALL
			SELECT 'color', id, name, hex, '' FROM colors
		),
		stats AS (
			SELECT
				MIN(price)::text AS min_price,
				MAX(price)::text AS max_price,
				MIN(width)::text AS min_width,
				MAX(width)::text AS max_width,
				MIN(height)::text AS min_height,
				MAX(height)::text AS max_height,
				MIN(depth)::text AS min_depth,
				MAX(depth)::text AS max_depth
			FROM products
			WHERE category_id = $1
		)
		SELECT * FROM dicts
		UNION ALL
		SELECT 'stats', '', min_price, max_price, min_width || ',' || max_width || ',' || min_height || ',' || max_height || ',' || min_depth || ',' || max_depth FROM stats;
	`

	rows, err := d.Driver.QueryContext(ctx, query, id)
	if err != nil {
		return nil, format.Error(op, err)
	}
	defer rows.Close()

	result := &productsRPC.Dictionaries{}
	for rows.Next() {
		var typ, id, field1, field2, field3 string
		if err := rows.Scan(&typ, &id, &field1, &field2, &field3); err != nil {
			log.Println(format.Error(op, err))
			continue
		}

		switch typ {
		case "brand":
			result.Brands.Brands = append(result.Brands.Brands, &productsRPC.Brand{Id: id, Title: field1})
		case "category":
			result.Categories.Categories = append(result.Categories.Categories, &productsRPC.Category{Id: id, Title: field1, Uri: field2})
		case "country":
			result.Countries.Countries = append(result.Countries.Countries, &productsRPC.Country{Id: id, Title: field1, Friendly: field2})
		case "material":
			result.Materials.Materials = append(result.Materials.Materials, &productsRPC.Material{Id: id, Title: field1})
		case "color":
			result.Colors.Colors = append(result.Colors.Colors, &productsRPC.Color{Id: id, Title: field1, Hex: field2})
		case "stats":
			// field1 = min_price, field2 = max_price, field3 = "minW,maxW,minH,maxH,minD,maxD"
			if minPrice, err := strconv.Atoi(field1); err == nil {
				result.MinPrice = int32(minPrice)
			}
			if maxPrice, err := strconv.Atoi(field2); err == nil {
				result.MaxPrice = int32(maxPrice)
			}
			parts := strings.Split(field3, ",")
			if len(parts) == 6 {
				r, _ := strconv.Atoi(parts[0])
				result.MinWidth = int32(r)
				r, _ = strconv.Atoi(parts[1])
				result.MaxWidth = int32(r)
				r, _ = strconv.Atoi(parts[2])
				result.MinHeight = int32(r)
				r, _ = strconv.Atoi(parts[3])
				result.MaxHeight = int32(r)
				r, _ = strconv.Atoi(parts[4])
				result.MinDepth = int32(r)
				r, _ = strconv.Atoi(parts[5])
				result.MaxDepth = int32(r)
			}
		}
	}

	return result, nil
}

func (d Driver) GetDictionaries(ctx context.Context) (*productsRPC.Dictionaries, error) {
	const op = "PostgresDb.GetDictionaries"

	query := `
		WITH dicts AS (
			SELECT 'brand' as type, id, name, '' as extra1, '' as extra2 FROM brands
			UNION ALL
			SELECT 'category', id, title, uri, '' FROM categories
			UNION ALL
			SELECT 'country', id, title, friendly, '' FROM countries
			UNION ALL
			SELECT 'material', id, title, '', '' FROM materials
			UNION ALL
			SELECT 'color', id, name, hex, '' FROM colors
		),
		stats AS (
			SELECT
				MIN(price)::text AS min_price,
				MAX(price)::text AS max_price,
				MIN(width)::text AS min_width,
				MAX(width)::text AS max_width,
				MIN(height)::text AS min_height,
				MAX(height)::text AS max_height,
				MIN(depth)::text AS min_depth,
				MAX(depth)::text AS max_depth
			FROM products
		)
		SELECT * FROM dicts
		UNION ALL
		SELECT 'stats', '', min_price, max_price, min_width || ',' || max_width || ',' || min_height || ',' || max_height || ',' || min_depth || ',' || max_depth FROM stats;
	`

	rows, err := d.Driver.QueryContext(ctx, query)
	if err != nil {
		return nil, format.Error(op, err)
	}
	defer rows.Close()

	result := &productsRPC.Dictionaries{}
	for rows.Next() {
		var typ, id, field1, field2, field3 string
		if err := rows.Scan(&typ, &id, &field1, &field2, &field3); err != nil {
			log.Println(format.Error(op, err))
			continue
		}

		switch typ {
		case "brand":
			result.Brands.Brands = append(result.Brands.Brands, &productsRPC.Brand{Id: id, Title: field1})
		case "category":
			result.Categories.Categories = append(result.Categories.Categories, &productsRPC.Category{Id: id, Title: field1, Uri: field2})
		case "country":
			result.Countries.Countries = append(result.Countries.Countries, &productsRPC.Country{Id: id, Title: field1, Friendly: field2})
		case "material":
			result.Materials.Materials = append(result.Materials.Materials, &productsRPC.Material{Id: id, Title: field1})
		case "color":
			result.Colors.Colors = append(result.Colors.Colors, &productsRPC.Color{Id: id, Title: field1, Hex: field2})
		case "stats":
			// field1 = min_price, field2 = max_price, field3 = "minW,maxW,minH,maxH,minD,maxD"
			if minPrice, err := strconv.Atoi(field1); err == nil {
				result.MinPrice = int32(minPrice)
			}
			if maxPrice, err := strconv.Atoi(field2); err == nil {
				result.MaxPrice = int32(maxPrice)
			}
			parts := strings.Split(field3, ",")
			if len(parts) == 6 {
				r, _ := strconv.Atoi(parts[0])
				result.MinWidth = int32(r)
				r, _ = strconv.Atoi(parts[1])
				result.MaxWidth = int32(r)
				r, _ = strconv.Atoi(parts[2])
				result.MinHeight = int32(r)
				r, _ = strconv.Atoi(parts[3])
				result.MaxHeight = int32(r)
				r, _ = strconv.Atoi(parts[4])
				result.MinDepth = int32(r)
				r, _ = strconv.Atoi(parts[5])
				result.MaxDepth = int32(r)
			}
		}
	}

	return result, nil
}
