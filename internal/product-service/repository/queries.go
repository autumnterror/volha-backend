package repository

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
	p.views,
	p.is_favorite,

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
		p.views,
		p.is_favorite,

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
					width, height, depth, photos, price, description, views, is_favorite
				)
				VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13, $14)
			),
			ins_materials AS (
				INSERT INTO product_materials (product_id, material_id)
				SELECT $1, unnest($15::varchar[])
			),
			ins_colors AS (
				INSERT INTO product_colors (product_id, color_id)
				SELECT $1, unnest($16::varchar[])
			),
			ins_seems AS (
				INSERT INTO product_seems (product_id, similar_product_id)
				SELECT $1, unnest($17::varchar[])
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
					description = $12,
					views 	    = $13,
					is_favorite = $14
				WHERE id = $1
			),
		
			-- materials
			ins_materials AS (
				INSERT INTO product_materials (product_id, material_id)
				SELECT $1, m_id
				FROM unnest($15::varchar[]) AS t(m_id)
				ON CONFLICT (product_id, material_id) DO NOTHING
			),
			del_materials AS (
				DELETE FROM product_materials pm
				WHERE pm.product_id = $1
				  AND NOT (pm.material_id = ANY($15::varchar[]))
			),
		
			-- colors
			ins_colors AS (
				INSERT INTO product_colors (product_id, color_id)
				SELECT $1, c_id
				FROM unnest($16::varchar[]) AS t(c_id)
				ON CONFLICT (product_id, color_id) DO NOTHING
			),
			del_colors AS (
				DELETE FROM product_colors pc
				WHERE pc.product_id = $1
				  AND NOT (pc.color_id = ANY($16::varchar[]))
			),
		
			-- seems
			ins_seems AS (
				INSERT INTO product_seems (product_id, similar_product_id)
				SELECT $1, s_id
				FROM unnest($17::varchar[]) AS t(s_id)
				ON CONFLICT (product_id, similar_product_id) DO NOTHING
			),
			del_seems AS (
				DELETE FROM product_seems ps
				WHERE ps.product_id = $1
				  AND NOT (ps.similar_product_id = ANY($17::varchar[]))
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
	p.views,
	p.is_favorite,

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
-- Выбираем бренды, которые есть у товаров в данной категории
SELECT DISTINCT 'brand' as type, b.id, b.title, '' as extra1, '' as extra2
FROM brands b
JOIN products p ON b.id = p.brand_id
WHERE p.category_id = $1

UNION ALL

-- Возвращаем все категории, чтобы пользователь мог переключаться между ними
SELECT 'category', id, title, uri, '' FROM categories

UNION ALL

-- Выбираем страны, которые есть у товаров в данной категории
SELECT DISTINCT 'country' as type, c.id, c.title, c.friendly, ''
FROM countries c
JOIN products p ON c.id = p.country_id
WHERE p.category_id = $1

UNION ALL

-- Выбираем материалы, которые есть у товаров в данной категории, через связующую таблицу
SELECT DISTINCT 'material' as type, m.id, m.title, '', ''
FROM materials m
JOIN product_materials pm ON m.id = pm.material_id
JOIN products p ON pm.product_id = p.id
WHERE p.category_id = $1

UNION ALL

-- Выбираем цвета, которые есть у товаров в данной категории, через связующую таблицу
SELECT DISTINCT 'color' as type, c.id, c.title, c.hex, ''
FROM colors c
JOIN product_colors pc ON c.id = pc.color_id
JOIN products p ON pc.product_id = p.id
WHERE p.category_id = $1

UNION ALL

-- Считаем статистику (мин/макс цены и размеры) только для товаров из этой категории
SELECT
    'stats',
    '',
    COALESCE(MIN(price)::text, '0'),
    COALESCE(MAX(price)::text, '0'),
    COALESCE(MIN(width)::text, '0') || ',' || COALESCE(MAX(width)::text, '0') || ',' ||
    COALESCE(MIN(height)::text, '0') || ',' || COALESCE(MAX(height)::text, '0') || ',' ||
    COALESCE(MIN(depth)::text, '0') || ',' || COALESCE(MAX(depth)::text, '0')
FROM products
WHERE category_id = $1;
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
