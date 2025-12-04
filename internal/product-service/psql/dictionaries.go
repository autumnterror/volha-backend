package psql

import (
	"context"
	"log"
	"strconv"
	"strings"

	"github.com/autumnterror/breezynotes/pkg/utils/format"
	productsRPC "github.com/autumnterror/volha-backend/pkg/proto/gen"
)

const NotByCategory = "CLR"

// GetDictionaries if field = NotByCategory then get all dict
func (d Driver) GetDictionaries(ctx context.Context, idCat string) (*productsRPC.Dictionaries, error) {
	const op = "PostgresDb.GetDictionaries"

	var query string
	var args []any
	if idCat == NotByCategory || idCat == "" {
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

	result := &productsRPC.Dictionaries{
		Brands:     &productsRPC.BrandList{Items: []*productsRPC.Brand{}},
		Categories: &productsRPC.CategoryList{Items: []*productsRPC.Category{}},
		Countries:  &productsRPC.CountryList{Items: []*productsRPC.Country{}},
		Materials:  &productsRPC.MaterialList{Items: []*productsRPC.Material{}},
		Colors:     &productsRPC.ColorList{Items: []*productsRPC.Color{}},
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
			result.Brands.Items = append(result.Brands.Items, &productsRPC.Brand{Id: id, Title: field1})
		case "category":
			result.Categories.Items = append(result.Categories.Items, &productsRPC.Category{Id: id, Title: field1, Uri: field2})
		case "country":
			result.Countries.Items = append(result.Countries.Items, &productsRPC.Country{Id: id, Title: field1, Friendly: field2})
		case "material":
			result.Materials.Items = append(result.Materials.Items, &productsRPC.Material{Id: id, Title: field1})
		case "color":
			result.Colors.Items = append(result.Colors.Items, &productsRPC.Color{Id: id, Title: field1, Hex: field2})
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
