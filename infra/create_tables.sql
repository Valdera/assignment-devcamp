CREATE TABLE IF NOT EXISTS products (
    id bigserial PRIMARY KEY,
    name varchar NOT NULL,
    price decimal NOT NULL,
    description text NOT NULL,
    variant varchar NOT NULL,
    discount decimal NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
);
