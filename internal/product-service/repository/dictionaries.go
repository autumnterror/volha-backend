package repository

import (
	"context"

	"log"
	"strconv"
	"strings"

	"github.com/autumnterror/breezynotes/pkg/utils/format"
	"github.com/autumnterror/volha-backend/internal/product-service/domain"
)

type DictRepo interface {
	GetDictionaries(ctx context.Context, idCat string) (*domain.Dictionaries, error)
}

// GetDictionaries if field = domain.NotByCategory then get all dict
func (d Driver) GetDictionaries(ctx context.Context, idCat string) (*domain.Dictionaries, error) {
	const op = "PostgresDb.GetDictionaries"

	var query string
	var args []any
	if idCat == domain.NotByCategory || idCat == "" {
		query = getDicQuery
	} else {
		query = getDicByCatQuery
		args = append(args, idCat)
	}

	rows, err := d.Driver.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, format.Error(op, err)
	}
	defer rows.Close()

	result := &domain.Dictionaries{
		Brands:     []*domain.Brand{},
		Categories: []*domain.Category{},
		Countries:  []*domain.Country{},
		Materials:  []*domain.Material{},
		Colors:     []*domain.Color{},
		MinPrice:   0,
		MaxPrice:   0,
		MinWidth:   0,
		MaxWidth:   0,
		MinHeight:  0,
		MaxHeight:  0,
		MinDepth:   0,
		MaxDepth:   0,
	}
	for rows.Next() {
		var typ, id, field1, field2, field3 string
		if err := rows.Scan(&typ, &id, &field1, &field2, &field3); err != nil {
			log.Println(format.Error(op, err))
			continue
		}

		switch typ {
		case "brand":
			result.Brands = append(result.Brands, &domain.Brand{Id: id, Title: field1})
		case "category":
			result.Categories = append(result.Categories, &domain.Category{Id: id, Title: field1, Uri: field2})
		case "country":
			result.Countries = append(result.Countries, &domain.Country{Id: id, Title: field1, Friendly: field2})
		case "material":
			result.Materials = append(result.Materials, &domain.Material{Id: id, Title: field1})
		case "color":
			result.Colors = append(result.Colors, &domain.Color{Id: id, Title: field1, Hex: field2})
		case "stats":
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
