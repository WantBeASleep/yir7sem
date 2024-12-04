-- +goose Up
-- +goose StatementBegin
CREATE TABLE device
(
    id   integer      PRIMARY KEY,
    name varchar(255) NOT NULL
);

COMMENT ON TABLE device IS 'Хранилище узи аппаратов на которых делались снимки';
COMMENT ON COLUMN device.name IS 'Название аппарата';

CREATE TABLE uzi
(
    id         uuid         PRIMARY KEY,
    projection varchar(255) NOT NULL,
    checked    boolean      NOT NULL,
    create_at  date         NOT NULL,
    patient_id uuid         NOT NULL,
    device_id  integer      NOT NULL REFERENCES device (id)
);

COMMENT ON TABLE uzi IS 'Хранилище описаний и характеристик узи';
COMMENT ON COLUMN uzi.projection IS 'Проекция в которой было сделано узи';
COMMENT ON COLUMN uzi.patient_id IS 'Идентификатор пациента к которому относится узи';
COMMENT ON COLUMN uzi.device_id IS 'Идентификатор узи аппарата на котором снято узи';

CREATE TABLE image
(
    id     uuid PRIMARY KEY,
    uzi_id uuid    NOT NULL REFERENCES uzi (id),
    page   integer NOT NULL
);

COMMENT ON TABLE image IS 'Хранилище кадров в узи';

CREATE TABLE node
(
    id        uuid    PRIMARY KEY,
    ai        boolean NOT NULL,
    tirads_23 real    NOT NULL,
    tirads_4  real    NOT NULL,
    tirads_5  real    NOT NULL
);

COMMENT ON TABLE node IS 'Хранилище узлов в узи';
COMMENT ON COLUMN node.ai IS 'Автор узла(нейронка ли)';
COMMENT ON COLUMN node.tirads_23 IS 'процент отношения к классу tirads_23';
COMMENT ON COLUMN node.tirads_4 IS 'процент отношения к классу tirads_4';
COMMENT ON COLUMN node.tirads_5 IS 'процент отношения к классу tirads_5';

CREATE TABLE segment
(
    id        uuid PRIMARY KEY,
    node_id   uuid NOT NULL REFERENCES node (id),
    image_id  uuid NOT NULL REFERENCES image (id),
    contor    text NOT NULL,
    tirads_23 real NOT NULL,
    tirads_4  real NOT NULL,
    tirads_5  real NOT NULL
);

COMMENT ON TABLE segment IS 'Хранилище сегментов в узи';
COMMENT ON COLUMN segment.contor IS 'контур узла (JSON)';
COMMENT ON COLUMN segment.tirads_23 IS 'процент отношения к классу tirads_23';
COMMENT ON COLUMN segment.tirads_4 IS 'процент отношения к классу tirads_4';
COMMENT ON COLUMN segment.tirads_5 IS 'процент отношения к классу tirads_5';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS node CASCADE;
DROP TABLE IF EXISTS segment CASCADE;
DROP TABLE IF EXISTS image CASCADE;
DROP TABLE IF EXISTS device CASCADE;
DROP TABLE IF EXISTS uzi CASCADE;
-- +goose StatementEnd
