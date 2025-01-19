-- +goose Up
-- +goose StatementBegin
ALTER TABLE node ADD uzi_id uuid REFERENCES uzi (id);

UPDATE node 
SET uzi_id = subq.uzi_id
FROM (
    SELECT DISTINCT i.uzi_id, n.id
    FROM image i
    JOIN segment s ON s.image_id = i.id
    JOIN node n ON s.node_id = n.id
) as subq
WHERE node.id = subq.id;

ALTER TABLE node ALTER COLUMN uzi_id SET NOT NULL;

COMMENT ON COLUMN node.uzi_id IS 'Идентификатор узи';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE node
    DROP COLUMN uzi_id;
-- +goose StatementEnd
