CREATE TABLE product_color_photos (
    product_id VARCHAR(100) NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    color_id VARCHAR(100) NOT NULL REFERENCES colors(id), photos VARCHAR(255)[] NOT NULL,
    PRIMARY KEY (product_id, color_id)
);