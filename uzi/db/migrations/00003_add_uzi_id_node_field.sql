-- +goose Up
-- +goose StatementBegin
ALTER TABLE node
    ADD uzi_id uuid NOT NULL REFERENCES uzi (id);

COMMENT ON COLUMN node.uzi_id IS 'Идентификатор узи';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE node
    DROP COLUMN uzi_id;
-- +goose StatementEnd
