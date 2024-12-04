-- +goose Up
-- +goose StatementBegin
CREATE TABLE "user"
(
    id       uuid         PRIMARY KEY,
    email    varchar(255) NOT NULL UNIQUE,
    password varchar(255) NOT NULL,
    token    text
);

COMMENT ON TABLE "user" IS 'Таблица пользователей';
COMMENT ON COLUMN "user".token IS 'Refresh токен пользователя';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "user" CASCADE;
-- +goose StatementEnd