CREATE TABLE inventories (
    id BIGSERIAL PRIMARY KEY,

    item_id VARCHAR(100) NOT NULL UNIQUE,

    item_name VARCHAR(255) NOT NULL,

    total_stock INTEGER NOT NULL CHECK (total_stock >= 0),

    reserved_stock INTEGER NOT NULL DEFAULT 0 CHECK (reserved_stock >= 0),

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_inventories_item_id
ON inventories(item_id);