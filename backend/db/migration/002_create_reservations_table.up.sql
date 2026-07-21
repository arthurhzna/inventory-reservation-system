CREATE TYPE reservation_status AS ENUM (
    'ACTIVE',
    'CONFIRMED',
    'EXPIRED'
);

CREATE TABLE reservations (

    id BIGSERIAL PRIMARY KEY,

    reservation_id UUID NOT NULL UNIQUE,

    inventory_id BIGINT NOT NULL,

    user_id VARCHAR(100) NOT NULL,

    quantity INTEGER NOT NULL CHECK (quantity > 0),

    status reservation_status NOT NULL,

    expires_at TIMESTAMP NOT NULL,

    confirmed_at TIMESTAMP,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_inventory
        FOREIGN KEY (inventory_id)
        REFERENCES inventories(id)
        ON DELETE CASCADE
);

CREATE UNIQUE INDEX idx_reservation_uuid
ON reservations(reservation_id);

CREATE INDEX idx_reservation_status
ON reservations(status);

CREATE INDEX idx_reservation_expired
ON reservations(status, expires_at);

CREATE INDEX idx_reservation_inventory
ON reservations(inventory_id);