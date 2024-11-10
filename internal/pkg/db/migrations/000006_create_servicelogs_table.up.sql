CREATE TABLE IF NOT EXISTS servicelogs (
    id UUID DEFAULT gen_random_uuid (),
    date date NOT NULL,
    mileage integer NOT NULL,
    description varchar NOT NULL,
    venue_id UUID NOT NULL,
    labor_cost float NOT NULL,
    created_at timestamp NOT NULL DEFAULT (now()),
    updated_at timestamp NOT NULL DEFAULT (now()),
    deleted_at timestamp,
    PRIMARY KEY (id)
);