-- +goose Up
-- +goose StatementBegin
CREATE TABLE echographic
(
    id                uuid PRIMARY KEY,
    contors           varchar(255),
    left_lobe_length  real,
    left_lobe_width   real,
    left_lobe_thick   real,
    left_lobe_volum   real,
    right_lobe_length real,
    right_lobe_width  real,
    right_lobe_thick  real,
    right_lobe_volum  real,
    gland_volum       real,
    isthmus           real,
    struct            varchar(255),
    echogenicity      varchar(255),
    regional_lymph    varchar(255),
    vascularization   varchar(255),
    location          varchar(255),
    additional        varchar(255),
    conclusion        varchar(255)
);

COMMENT ON TABLE echographic IS 'Эхографические признаки';
COMMENT ON COLUMN echographic.id IS 'ID uzi';
COMMENT ON COLUMN echographic.contors IS 'Контуры';
COMMENT ON COLUMN echographic.left_lobe_length IS 'Левая доля длина';
COMMENT ON COLUMN echographic.left_lobe_width IS 'Левая доля ширина';
COMMENT ON COLUMN echographic.left_lobe_thick IS 'Левая доля толщина';
COMMENT ON COLUMN echographic.left_lobe_volum IS 'Левая доля объем';
COMMENT ON COLUMN echographic.right_lobe_length IS 'Правая доля длина';
COMMENT ON COLUMN echographic.right_lobe_width IS 'Правая доля ширина';
COMMENT ON COLUMN echographic.right_lobe_thick IS 'Правая доля толщина';
COMMENT ON COLUMN echographic.right_lobe_volum IS 'Правая доля объем';
COMMENT ON COLUMN echographic.gland_volum IS 'Объем железы';
COMMENT ON COLUMN echographic.isthmus IS 'Перешеек';
COMMENT ON COLUMN echographic.struct IS 'Структура';
COMMENT ON COLUMN echographic.echogenicity IS 'Эхогенность';
COMMENT ON COLUMN echographic.regional_lymph IS 'Региональные лимфатические узлы';
COMMENT ON COLUMN echographic.vascularization IS 'Васкуляризацие по ЦДК';
COMMENT ON COLUMN echographic.location IS 'Расположение';
COMMENT ON COLUMN echographic.additional IS 'Дополнительные данные';
COMMENT ON COLUMN echographic.conclusion IS 'Заключение';

ALTER TABLE echographic
    ADD CONSTRAINT fk_echographic_uzi
    FOREIGN KEY (id)
    REFERENCES uzi (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE echographic
    DROP CONSTRAINT fk_echographic_uzi;

DROP TABLE IF EXISTS echographic CASCADE;
-- +goose StatementEnd
