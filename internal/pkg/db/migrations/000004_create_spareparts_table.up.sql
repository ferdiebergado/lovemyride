CREATE TABLE IF NOT EXISTS spareparts (
    id UUID DEFAULT gen_random_uuid (),
    description varchar NOT NULL,
    maintenance_interval integer NOT NULL,
    created_at timestamp NOT NULL DEFAULT (now()),
    updated_at timestamp NOT NULL DEFAULT (now()),
    deleted_at timestamp,
    PRIMARY KEY (id)
);