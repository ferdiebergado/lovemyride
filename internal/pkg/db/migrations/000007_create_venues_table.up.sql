CREATE TABLE IF NOT EXISTS venues (
    id UUID DEFAULT gen_random_uuid (),
    name varchar NOT NULL,
    created_at timestamp NOT NULL DEFAULT (now()),
    updated_at timestamp NOT NULL DEFAULT (now()),
    deleted_at timestamp,
    PRIMARY KEY (id)
);