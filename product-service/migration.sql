
CREATE TABLE products (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    sku TEXT NOT NULL UNIQUE,
    description TEXT,
    price NUMERIC(10, 2) NOT NULL,
    image_url TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);


CREATE INDEX idx_products_name ON products (name);
CREATE UNIQUE INDEX idx_products_sku ON products (sku);
CREATE INDEX idx_products_deleted_at ON products (deleted_at);