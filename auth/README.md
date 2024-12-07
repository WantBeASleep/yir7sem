# Auth сервис

## Reqire
_лучше подгружать через `.env`_
| Env | Value | Описание |
|----------|----------|----------|
|GOOSE_DRIVER| postgres | `env` для миграций |
|GOOSE_DBSTRING| postgres dsn  | dsn postres sql для миграций |
| JWT_KEY_PRIVATE   | `rsa256`   | ключ шифрования   |
| JWT_KEY_PUBLIC    | `rsa256`   | ключ шифрования   |
|JWT_TOKEN_AC_TIME| время (5m) | время действия access токена |
|JWT_TOKEN_RE_TIME| время (24h) | время действия refresh токена |
|DB_DSN| postgres dsn | dsn для postgres sql |
|APP_URL| localhost:50055 | grpc url |



Основная цель сервиса - обеспечить пользователя jwt токеном, для аунтентификации и авторизации.
_В будующем можно накрутить сюда систему ролей_

Используется _jwt_, шифрования по схеме `RS256`
jwt : header, payload, signature
private key: header, payload --> signature
public key: signature --> header, payload

таким образом только наш сервис умеет создавать ключи, остальные лиш проверять что они были выданны им
//TODO: починить ошибки если приходят невалидные данные