CREATE TABLE IF NOT EXISTS spareparts_prices (
    sparepart_id UUID NOT NULL,
    price float NOT NULL,
    purchase_date date NOT NULL,
    store_id UUID NOT NULL,
    created_at timestamp NOT NULL DEFAULT (now()),
    updated_at timestamp NOT NULL DEFAULT (now()),
    deleted_at timestamp,
    PRIMARY KEY (sparepart_id)
);