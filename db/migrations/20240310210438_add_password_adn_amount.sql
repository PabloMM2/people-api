-- migrate:up
ALTER TABLE person
ADD COLUMN password VARCHAR(255),
ADD COLUMN amount DOUBLE PRECISION;

-- migrate:down

ALTER TABLE person
DROP COLUMN IF EXISTS password,
DROP COLUMN IF EXISTS amount;
