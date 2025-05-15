CREATE TABLE orders (
    uuid            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id        TEXT NOT NULL,
    user_uuid       UUID NOT NULL,
    status          TEXT NOT NULL,
    total_quantity  INTEGER NOT NULL DEFAULT 0,
    total_items     INTEGER NOT NULL DEFAULT 0,
    created_at      TIMESTAMPTZ DEFAULT NOW(),
    updated_at      TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE order_items (
    uuid            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_uuid      UUID NOT NULL,
    product_uuid    UUID NOT NULL,
    store_uuid      UUID,
    warehouse_uuid  UUID,
    product_name    TEXT NOT NULL,
    product_sku     TEXT NOT NULL,
    quantity        INTEGER NOT NULL,
    price           NUMERIC(12, 2) NOT NULL,

    CONSTRAINT fk_order FOREIGN KEY (order_uuid) REFERENCES orders(uuid)
);

CREATE INDEX idx_order_items_product_uuid ON order_items (product_uuid);
CREATE INDEX idx_order_items_warehouse_uuid ON order_items (warehouse_uuid);

