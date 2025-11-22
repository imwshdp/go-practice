CREATE TYPE order_status AS ENUM ('pending', 'completed', 'cancelled');

CREATE TABLE IF NOT EXISTS orders (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  total NUMERIC(10, 2) NOT NULL,
  status order_status NOT NULL DEFAULT 'pending',
  address TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

  CONSTRAINT fk_orders_user
    FOREIGN KEY (user_id) REFERENCES users(id)
    ON DELETE CASCADE
);
