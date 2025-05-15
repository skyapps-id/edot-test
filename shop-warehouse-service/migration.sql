CREATE TABLE shops (
    uuid         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name         TEXT NOT NULL,
    address      TEXT,
    created_at   TIMESTAMPTZ DEFAULT NOW(),
    updated_at   TIMESTAMPTZ DEFAULT NOW(),
    deleted_at   TIMESTAMPTZ
);

CREATE TABLE warehouses (
    uuid         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name         TEXT NOT NULL,
    address      TEXT,
    active       BOOLEAN DEFAULT TRUE,
    shop_uuid    UUID NOT NULL,
    created_at   TIMESTAMPTZ DEFAULT NOW(),
    updated_at   TIMESTAMPTZ DEFAULT NOW(),
    deleted_at   TIMESTAMPTZ,

    CONSTRAINT fk_shop FOREIGN KEY (shop_uuid) REFERENCES shops(uuid)
);

CREATE TABLE warehouse_products (
    uuid             UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    warehouse_uuid   UUID NOT NULL,
    product_uuid     UUID NOT NULL,
    quantity         INTEGER NOT NULL DEFAULT 0,
    updated_at       TIMESTAMPTZ DEFAULT NOW(),

    CONSTRAINT fk_warehouse FOREIGN KEY (warehouse_uuid) REFERENCES warehouses(uuid)
);

CREATE INDEX idx_shops_name ON shops(name);
CREATE INDEX idx_warehouses_name ON warehouses(name);
CREATE INDEX idx_warehouses_shop_uuid ON warehouses(shop_uuid);
CREATE INDEX idx_warehouse_products_warehouse_uuid ON warehouse_products(warehouse_uuid);