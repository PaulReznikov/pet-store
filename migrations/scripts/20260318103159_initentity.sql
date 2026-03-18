-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS brands (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  image_url TEXT
);
CREATE TABLE IF NOT EXISTS animals (
  id SERIAL PRIMARY KEY,
  type VARCHAR(255) NOT NULL,
  image_url TEXT
);
CREATE TABLE IF NOT EXISTS categories (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  parent_id INT REFERENCES categories(id) ON DELETE RESTRICT
);
CREATE TABLE IF NOT EXISTS sales (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  discount_percentage DECIMAL(5, 2) NOT NULL,
  start_date TIMESTAMP WITH TIME ZONE NOT NULL,
  end_date TIMESTAMP WITH TIME ZONE NOT NULL
);
CREATE TABLE IF NOT EXISTS products (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  img_url TEXT,
  description TEXT,
  price DECIMAL(10, 2) NOT NULL,
  category_id INT REFERENCES categories(id) ON DELETE RESTRICT,
  sale_id INT REFERENCES sales(id) ON DELETE SET NULL,
  brand_id INT REFERENCES brands(id) ON DELETE RESTRICT,
  animal_id INT REFERENCES animals(id) ON DELETE RESTRICT,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS product_packages (
  id SERIAL PRIMARY KEY,
  value DECIMAL(10,2) NOT NULL,
  unit VARCHAR(50) NOT NULL,
  count INT NOT NULL,
  price DECIMAL(10, 2) NOT NULL,
  product_id INT REFERENCES products(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS product_packages;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS sales;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS brands;
DROP TABLE IF EXISTS animals;
-- +goose StatementEnd
