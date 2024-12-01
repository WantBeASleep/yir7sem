-- +goose Up
-- +goose StatementBegin
CREATE TABLE doctor
(
    id       uuid PRIMARY KEY,
    fullname varchar(255) NOT NULL,
    org      varchar(255) NOT NULL,
    job      varchar(255) NOT NULL,
    "desc"   text
);

COMMENT ON TABLE doctor IS 'Таблица врачей';
COMMENT ON COLUMN doctor.id IS 'Идентификатор, совпадает с ID пользователя из auth';
COMMENT ON COLUMN doctor.org IS 'Мед организация доктора';
COMMENT ON COLUMN doctor.job IS 'Должность доктора';
COMMENT ON COLUMN doctor.desc IS 'Описание, опыт работы врача';

CREATE TABLE patient
(
    id            uuid PRIMARY KEY,
    fullName      varchar(255) NOT NULL,
    email         varchar(255) NOT NULL,
    policy        varchar(255) NOT NULL,
    active        boolean      NOT NULL,
    malignancy    boolean      NOT NULL,
    last_uzi_date timestamp
);

COMMENT ON TABLE patient IS 'Таблица пациентов';
COMMENT ON COLUMN patient.policy IS 'Мед полис';
COMMENT ON COLUMN patient.active IS 'Активен ли пациент';
COMMENT ON COLUMN patient.malignancy IS 'Показатель злокачественного образования';
COMMENT ON COLUMN patient.last_uzi_date IS 'Время последенго узи снимка';

CREATE TABLE card
(
    doctor_id  uuid,
    patient_id uuid,
    diagnosis  text
);

ALTER TABLE card
    ADD CONSTRAINT pk_card PRIMARY KEY (doctor_id, patient_id);

COMMENT ON TABLE card IS 'Таблица мед карт';
COMMENT ON COLUMN card.diagnosis IS 'Диагноз конкретного врача по пациенту';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS card CASCADE;
DROP TABLE IF EXISTS doctor CASCADE;
DROP TABLE IF EXISTS patient CASCADE;
-- +goose StatementEnd
