-- +goose Up
ALTER TABLE patient ADD COLUMN birthdate TIMESTAMP;

-- +goose Down
ALTER TABLE patient DROP COLUMN birthdate; 