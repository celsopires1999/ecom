CREATE TABLE IF NOT EXISTS order_items (
  order_item_id SERIAL PRIMARY KEY,
  order_id INTEGER NOT NULL REFERENCES orders (order_id) ON DELETE RESTRICT,
  product_id INTEGER NOT NULL REFERENCES products (product_id) ON DELETE RESTRICT,
  quantity INTEGER NOT NULL,
  price NUMERIC(10, 2) NOT NULL
);
