-- +goose Up
ALTER TABLE patient ADD COLUMN birthdate TIMESTAMP NOT NULL;

-- +goose Down
ALTER TABLE patient DROP COLUMN birthdate; 