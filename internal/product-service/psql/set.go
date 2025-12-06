package psql

import (
	"context"
	"database/sql"
	"errors"
	productsRPC "github.com/autumnterror/volha-backend/pkg/proto/gen"
	"github.com/autumnterror/volha-backend/pkg/views"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrUnknownType   = errors.New("unknown type")
	ErrInvalidType   = errors.New("bad type of obj")
	ErrAlreadyExists = errors.New("obj already exist")
	ErrForeignKey    = errors.New("sub obj dont exist")
)

type SqlRepo interface {
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

type Driver struct {
	Driver SqlRepo
}

type Repo interface {
	FilterProducts(ctx context.Context, filter *productsRPC.ProductFilter) ([]*productsRPC.Product, error)
	SearchProducts(ctx context.Context, filter *productsRPC.ProductSearch) ([]*productsRPC.Product, error)

	GetAllProducts(ctx context.Context, start, end int) ([]*productsRPC.Product, error)
	GetProduct(ctx context.Context, id string) (*productsRPC.Product, error)
	CreateProduct(ctx context.Context, p *productsRPC.ProductId) error
	UpdateProduct(ctx context.Context, p *productsRPC.ProductId) error
	DeleteProduct(ctx context.Context, id string) error

	GetDictionaries(ctx context.Context, idCat string) (*productsRPC.Dictionaries, error)

	GetAll(ctx context.Context, _type views.Type) (any, error)
	Get(ctx context.Context, id string, _type views.Type) (any, error)
	Create(ctx context.Context, obj any, _type views.Type) error
	Update(ctx context.Context, obj any, _type views.Type) error
	Delete(ctx context.Context, id string, _type views.Type) error

	GetProductColorPhotos(ctx context.Context, productID string, colorID string) (any, error)
	DeleteProductColorPhotos(ctx context.Context, productID string, colorID string) error
}

var emptyProduct = productsRPC.Product{
	Id:      "",
	Title:   "",
	Article: "",
	Brand: &productsRPC.Brand{
		Id:    "",
		Title: "",
	},
	Category: &productsRPC.Category{
		Id:    "",
		Title: "",
		Uri:   "",
		Img:   "",
	},
	Country: &productsRPC.Country{
		Id:       "",
		Title:    "",
		Friendly: "",
	},
	Width:       0,
	Height:      0,
	Depth:       0,
	Materials:   []*productsRPC.Material{},
	Colors:      []*productsRPC.Color{},
	Photos:      []string{},
	Seems:       []*productsRPC.Product{},
	Price:       0,
	Description: "",
}

const getAllProductsQuery = `
SELECT
    p.id,
    p.title,
    p.article,
    p.width,
    p.height,
    p.depth,
    p.photos,
    p.price,
    p.description,

    b.id      AS brand_id,
    b.title   AS brand_title,

    cat.id    AS category_id,
    cat.title AS category_title,
    cat.uri   AS category_uri,

    co.id       AS country_id,
    co.title    AS country_title,
    co.friendly AS country_friendly,

    COALESCE(
        (
            SELECT json_agg(
                       json_build_object(
                           'id', m.id,
                           'title', m.title
                       )
                   )
            FROM product_materials pm
            JOIN materials m ON m.id = pm.material_id
            WHERE pm.product_id = p.id
        ),
        '[]'::json
    ) AS materials,

    COALESCE(
        (
            SELECT json_agg(
                       json_build_object(
                           'id', c.id,
                           'title', c.title,
                           'hex', c.hex
                       )
                   )
            FROM product_colors pc
            JOIN colors c ON c.id = pc.color_id
            WHERE pc.product_id = p.id
        ),
        '[]'::json
    ) AS colors,

    COALESCE(
        (
            SELECT json_agg(
                       json_build_object(
                           'id', sp.id,
                           'title', sp.title,
                           'article', sp.article,
                           'width', sp.width,
                           'height', sp.height,
                           'depth', sp.depth,
                           'photos', sp.photos,
                           'price', sp.price,
                           'description', sp.description
                       )
                   )
            FROM product_seems ps
            JOIN products sp ON sp.id = ps.similar_product_id
            WHERE ps.product_id = p.id
        ),
        '[]'::json
    ) AS similar_products

FROM products p
JOIN brands     b   ON b.id   = p.brand_id
JOIN categories cat ON cat.id = p.category_id
JOIN countries  co  ON co.id  = p.country_id
ORDER BY p.title
LIMIT $1 OFFSET $2
`
const getProductQuery = `
	SELECT
		p.id,
		p.title,
		p.article,
		p.width,
		p.height,
		p.depth,
		p.photos,
		p.price,
		p.description,
	
		b.id      AS brand_id,
		b.title   AS brand_title,
	
		cat.id    AS category_id,
		cat.title AS category_title,
		cat.uri   AS category_uri,
	
		co.id       AS country_id,
		co.title    AS country_title,
		co.friendly AS country_friendly,
	
		COALESCE(
			(
				SELECT json_agg(
						   json_build_object(
							   'id', m.id,
							   'title', m.title
						   )
					   )
				FROM product_materials pm
				JOIN materials m ON m.id = pm.material_id
				WHERE pm.product_id = p.id
			),
			'[]'::json
		) AS materials,
	
		COALESCE(
			(
				SELECT json_agg(
						   json_build_object(
							   'id', c.id,
							   'title', c.title,
							   'hex', c.hex
						   )
					   )
				FROM product_colors pc
				JOIN colors c ON c.id = pc.color_id
				WHERE pc.product_id = p.id
			),
			'[]'::json
		) AS colors,
	
		COALESCE(
			(
				SELECT json_agg(
						   json_build_object(
							   'id', sp.id,
							   'title', sp.title,
							   'article', sp.article,
							   'width', sp.width,
							   'height', sp.height,
							   'depth', sp.depth,
							   'photos', sp.photos,
							   'price', sp.price,
							   'description', sp.description
						   )
					   )
				FROM product_seems ps
				JOIN products sp ON sp.id = ps.similar_product_id
				WHERE ps.product_id = p.id
			),
			'[]'::json
		) AS similar_products
	
	FROM products p
	JOIN brands     b   ON b.id   = p.brand_id
	JOIN categories cat ON cat.id = p.category_id
	JOIN countries  co  ON co.id  = p.country_id
	WHERE p.id = $1
`
const createProductQuery = `
		WITH
			ins_product AS (
				INSERT INTO products (
					id, title, article, brand_id, category_id, country_id,
					width, height, depth, photos, price, description
				)
				VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
			),
			ins_materials AS (
				INSERT INTO product_materials (product_id, material_id)
				SELECT $1, unnest($13::varchar[])
			),
			ins_colors AS (
				INSERT INTO product_colors (product_id, color_id)
				SELECT $1, unnest($14::varchar[])
			),
			ins_seems AS (
				INSERT INTO product_seems (product_id, similar_product_id)
				SELECT $1, unnest($15::varchar[])
			)
		SELECT 1;
	`
const updateProductQuery = `
		WITH
			upd_product AS (
				UPDATE products SET 
					title       = $2,
					article     = $3,
					brand_id    = $4,
					category_id = $5,
					country_id  = $6,
					width       = $7,
					height      = $8,
					depth       = $9,
					photos      = $10,
					price       = $11,
					description = $12
				WHERE id = $1
			),
		
			-- materials
			ins_materials AS (
				INSERT INTO product_materials (product_id, material_id)
				SELECT $1, m_id
				FROM unnest($13::varchar[]) AS t(m_id)
				ON CONFLICT (product_id, material_id) DO NOTHING
			),
			del_materials AS (
				DELETE FROM product_materials pm
				WHERE pm.product_id = $1
				  AND NOT (pm.material_id = ANY($13::varchar[]))
			),
		
			-- colors
			ins_colors AS (
				INSERT INTO product_colors (product_id, color_id)
				SELECT $1, c_id
				FROM unnest($14::varchar[]) AS t(c_id)
				ON CONFLICT (product_id, color_id) DO NOTHING
			),
			del_colors AS (
				DELETE FROM product_colors pc
				WHERE pc.product_id = $1
				  AND NOT (pc.color_id = ANY($14::varchar[]))
			),
		
			-- seems
			ins_seems AS (
				INSERT INTO product_seems (product_id, similar_product_id)
				SELECT $1, s_id
				FROM unnest($15::varchar[]) AS t(s_id)
				ON CONFLICT (product_id, similar_product_id) DO NOTHING
			),
			del_seems AS (
				DELETE FROM product_seems ps
				WHERE ps.product_id = $1
				  AND NOT (ps.similar_product_id = ANY($15::varchar[]))
			)
		SELECT 1;

	`
const filterProductsQuery = `
SELECT
    p.id,
    p.title,
    p.article,
    p.width,
    p.height,
    p.depth,
    p.photos,
    p.price,
    p.description,

    b.id      AS brand_id,
    b.title   AS brand_title,

    cat.id    AS category_id,
    cat.title AS category_title,
    cat.uri   AS category_uri,

    co.id       AS country_id,
    co.title    AS country_title,
    co.friendly AS country_friendly,

    COALESCE(
        (
            SELECT json_agg(
                       json_build_object(
                           'id', m.id,
                           'title', m.title
                       )
                   )
            FROM product_materials pm
            JOIN materials m ON m.id = pm.material_id
            WHERE pm.product_id = p.id
        ),
        '[]'::json
    ) AS materials,

    COALESCE(
        (
            SELECT json_agg(
                       json_build_object(
                           'id', c.id,
                           'title', c.title,
                           'hex', c.hex
                       )
                   )
            FROM product_colors pc
            JOIN colors c ON c.id = pc.color_id
            WHERE pc.product_id = p.id
        ),
        '[]'::json
    ) AS colors,

    COALESCE(
        (
            SELECT json_agg(
                       json_build_object(
                           'id', sp.id,
                           'title', sp.title,
                           'article', sp.article,
                           'width', sp.width,
                           'height', sp.height,
                           'depth', sp.depth,
                           'photos', sp.photos,
                           'price', sp.price,
                           'description', sp.description
                       )
                   )
            FROM product_seems ps
            JOIN products sp ON sp.id = ps.similar_product_id
            WHERE ps.product_id = p.id
        ),
        '[]'::json
    ) AS similar_products

FROM products p
JOIN brands     b   ON b.id   = p.brand_id
JOIN categories cat ON cat.id = p.category_id
JOIN countries  co  ON co.id  = p.country_id
`
const getDicByCatQuery = `
		WITH dicts AS (
			SELECT 'brand' as type, id, title, '' as extra1, '' as extra2 FROM brands
			UNION ALL
			SELECT 'category', id, title, uri, '' FROM categories
			UNION ALL
			SELECT 'country', id, title, friendly, '' FROM countries
			UNION ALL
			SELECT 'material', id, title, '', '' FROM materials
			UNION ALL
			SELECT 'color', id, title, hex, '' FROM colors
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
const getDicQuery = `
		WITH dicts AS (
			SELECT 'brand' as type, id, title, '' as extra1, '' as extra2 FROM brands
			UNION ALL
			SELECT 'category', id, title, uri, '' FROM categories
			UNION ALL
			SELECT 'country', id, title, friendly, '' FROM countries
			UNION ALL
			SELECT 'material', id, title, '', '' FROM materials
			UNION ALL
			SELECT 'color', id, title, hex, '' FROM colors
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
