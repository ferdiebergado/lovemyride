CREATE TABLE IF NOT EXISTS servicelogs_spareparts (
    id UUID DEFAULT gen_random_uuid (),
    servicelog_id UUID NOT NULL,
    sparepart_id UUID NOT NULL,
    created_at timestamp NOT NULL DEFAULT (now()),
    updated_at timestamp NOT NULL DEFAULT (now()),
    deleted_at timestamp,
    PRIMARY KEY (id)
);