-- migrations/000001_init.up.sql

-- create shops table
CREATE TABLE IF NOT EXISTS shops (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  address VARCHAR(255),
  contact_phone VARCHAR(255),
  description TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT Now()
);

-- create grocery list table
CREATE TABLE IF NOT EXISTS grocery_lists (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  description TEXT,
  status VARCHAR(255) NOT NULL,
  completed_at TIMESTAMP,
  estimated_total DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
  actual_total DECIMAL(10, 2),
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT Now()
);

-- create grocery_items table
CREATE TABLE IF NOT EXISTS grocery_items (
  id SERIAL PRIMARY KEY,
  grocery_list_id INTEGER NOT NULL REFERENCES grocery_lists(id),
  name VARCHAR(255) NOT NULL,
  description TEXT,
  quantity DECIMAL(10, 3) NOT NULL DEFAULT 1,
  unit VARCHAR(255) NOT NULL,
  estimated_price DECIMAL(10, 2) NOT NULL,
  actual_price DECIMAL(10, 2),
  is_purchased BOOLEAN NOT NULL DEFAULT FALSE,
  shop_id INTEGER NOT NULL REFERENCES shops(id),
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_grocery_items_grocery_list_id ON grocery_items (grocery_list_id);
CREATE INDEX idx_grocery_items_shop_id ON grocery_items (shop_id);