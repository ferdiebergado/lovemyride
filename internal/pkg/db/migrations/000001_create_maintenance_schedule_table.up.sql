CREATE TABLE IF NOT EXISTS maintenance_schedule (
    id UUID DEFAULT gen_random_uuid (),
    sparepart_id UUID NOT NULL,
    next_mileage integer NOT NULL,
    created_at timestamp NOT NULL DEFAULT (now()),
    updated_at timestamp NOT NULL DEFAULT (now()),
    deleted_at timestamp,
    PRIMARY KEY (id)
);