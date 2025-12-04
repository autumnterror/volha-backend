-- Таблица брендов
CREATE TABLE brands (
                        id VARCHAR(100) PRIMARY KEY,
                        title VARCHAR(100) NOT NULL
);

-- Таблица категорий
CREATE TABLE categories (
                            id VARCHAR(100) PRIMARY KEY,
                            title VARCHAR(100) NOT NULL,
                            uri VARCHAR(100) NOT NULL
);

-- Таблица стран
CREATE TABLE countries (
                           id VARCHAR(100) PRIMARY KEY,
                           title VARCHAR(100) NOT NULL,
                           friendly VARCHAR(100)
);

-- Таблица материалов
CREATE TABLE materials (
                           id VARCHAR(100) PRIMARY KEY,
                           title VARCHAR(100) NOT NULL
);

-- Таблица цветов
CREATE TABLE colors (
                        id VARCHAR(100) PRIMARY KEY,
                        title VARCHAR(100) NOT NULL,
                        hex VARCHAR(7) NOT NULL
);

-- Основная таблица продуктов
CREATE TABLE products (
                          id VARCHAR(100) PRIMARY KEY,
                          title VARCHAR(100) NOT NULL UNIQUE,
                          article VARCHAR(100) NOT NULL UNIQUE,
                          brand_id VARCHAR(100) NOT NULL REFERENCES brands(id),
                          category_id VARCHAR(100) NOT NULL REFERENCES categories(id),
                          country_id VARCHAR(100) NOT NULL REFERENCES countries(id),
                          width INT NOT NULL,
                          height INT NOT NULL,
                          depth INT NOT NULL,
                          photos VARCHAR(255)[] NOT NULL,
                          price INT NOT NULL,
                          description TEXT NOT NULL
);

-- Таблица связей продуктов с материалами
CREATE TABLE product_materials (
                                   product_id VARCHAR(100) NOT NULL REFERENCES products(id) ON DELETE CASCADE,
                                   material_id VARCHAR(100) NOT NULL REFERENCES materials(id),
                                   PRIMARY KEY (product_id, material_id)
);

-- Таблица связей продуктов с цветами
CREATE TABLE product_colors (
                                product_id VARCHAR(100) NOT NULL REFERENCES products(id) ON DELETE CASCADE,
                                color_id VARCHAR(100) NOT NULL REFERENCES colors(id),
                                PRIMARY KEY (product_id, color_id)
);

-- Таблица связей "похожих товаров"
CREATE TABLE product_seems (
                               product_id VARCHAR(100) NOT NULL REFERENCES products(id) ON DELETE CASCADE,
                               similar_product_id VARCHAR(100),
                               PRIMARY KEY (product_id, similar_product_id)
);
