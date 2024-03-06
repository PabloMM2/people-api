-- migrate:up
CREATE TABLE IF NOT EXISTS public.person (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    last_name VARCHAR(255),
    email VARCHAR(255),
    age INTEGER
);

-- migrate:down
DROP TABLE IF EXISTS public.person;
