# gateway

Точка входа в систему, все вызовы проходят api-gateway
//TODO: is checked в узи расписать

## Reqire
_лучше подгружать через `.env`_
| Env | Value | Описание |
|----------|----------|----------|
|APP_URL| localhost:8080 | http url |
| JWT_KEY_PUBLIC    | `rsa256`   | ключ шифрования   |
|S3_ENDPOINT| localhost:9000 | s3 addrs |
|S3_TOKEN_ACCESS| key | pub ключ для S3 |
|S3_TOKEN_SECRET| secret key | priv ключ для S3 |
|BROKER_ADDRS| localhost:19092 | url для брокера (массив) |
|ADAPTERS_AUTHURL|localhost:50040|url auth сервиса|
|ADAPTERS_MEDURL|localhost:50050|url med сервиса|
|ADAPTERS_UZIURL|localhost:50060|url uzi сервиса|