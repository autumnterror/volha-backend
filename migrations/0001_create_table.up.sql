CREATE TABLE products
(
    id    VARCHAR(100) PRIMARY KEY,
    title VARCHAR(100),
    article VARCHAR(100) UNIQUE,
    brand VARCHAR(100),
    country VARCHAR(100),
    width int,
    height int,
    depth int,
    materials VARCHAR(100)[],
    color VARCHAR(100)[],
    photos VARCHAR(100)[],
    seems VARCHAR(100)[],
    price int,
    description text
);