-- +migrate Down
-- +migrate StatementBegin

DROP TABLE IF EXISTS books CASCADE;
DROP TABLE IF EXISTS categories CASCADE;
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS person CASCADE;

-- +migrate StatementEnd