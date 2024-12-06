# Auth сервис

## Reqire
_лучше подгружать через `.env`_
| Env | Value | Описание |
|----------|----------|----------|
|GOOSE_DRIVER| postgres | `env` для миграций |
|GOOSE_DBSTRING| postgres dsn  | dsn postres sql для миграций |
| JWT_KEY_PRIVATE   | `rsa256`   | ключ шифрования   |
| JWT_KEY_PUBLIC    | `rsa256`   | ключ шифрования   |
|DB_DSN| postgres dsn | dsn для postgres sql |



Основная цель сервиса - обеспечить пользователя jwt токеном, для аунтентификации и авторизации.
_В будующем можно накрутить сюда систему ролей_

Используется _jwt_, шифрования по схеме `RS256`
jwt : header, payload, signature
private key: header, payload --> signature
public key: signature --> header, payload

таким образом только наш сервис умеет создавать ключи, остальные лиш проверять что они были выданны им