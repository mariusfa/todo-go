CREATE SCHEMA IF NOT EXISTS todoschema;

CREATE USER appuser WITH PASSWORD 'password';

ALTER USER appuser SET search_path TO todoschema;

GRANT USAGE ON SCHEMA todoschema TO appuser;

CREATE TABLE IF NOT EXISTS todoschema.todos (
    id SERIAL PRIMARY KEY,
    task VARCHAR NOT NULL
);

GRANT SELECT, INSERT, UPDATE, DELETE ON todoschema.todos TO appuser;
GRANT USAGE, SELECT ON SEQUENCE todoschema.todos_id_seq TO appuser;
