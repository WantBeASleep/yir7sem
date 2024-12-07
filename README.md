# Сервис обработки узи МРТ снимков

## Общая архитектура

![all_services_shema](./docs/assets/all_services_shema.png)

Приведенная схема отражает список сервисов.

`docker`:
* 8080 - api-gateway
* 8080/docs - swagger
* 8081 - redpanda console (kafka, не накликайте ничего)
* 9000 - s3
    - log, pss: minioadmin:minioadmin
* 50050 - grpc auth service
* 50051 - grpc med service
* 50052 - grpc uzi service

_psql db: authdb, meddb, uzidb_. Создаются при инициализации docker volume. Миграции накатываются при запуске контейнера. __ПОРТ 5433__

__КАК ЗАПУСТИТЬ?__
0) добыть веса для нейронки (Г-52 Пн-Вт 13:00 - 13:15) и положить ml_service/internal/ml_model/models/cross и all
1) docker compose --profile deps up -d
2) создать бакет в минио + access + secret key
    + зайти на localhost:9000
    + minioadmin:minioadmin
    + слева вкладке access keys
    + справа сверху `create access key` - вбиваем в него
        - access_key: `NZjt6KmuHQRU7IitYUiW`
        - secret_key: `ql9DoBMKyqMxQm8j5LQuKwnn68KFsGqn5jGbL7uL`
3) вкладка object store, создаем бакет uzi
4) создаем топики для redpanda (потом автоматизируем, сейчас я в __`тильте`__)
    + localhost:8081
    + вкладка topics
    + создаем 3 топика:
        - uziupload
        - uzisplitted
        - uziprocessed
5) docker compose --profile app up -d

Если не сносить docker volume's, то операцию нужно будет делать всего 1 раз

_P.S.: хз как сделать это через env'арники, только какая то гига муть с доп контейнером для sh. Накликайте руками, надежднее будет. Ключи можете поменять, тогда не забудьте прокинуть их в сервисы через .env-docker файлы/запускать контейнеры с своими env'арами. В коде никакие ключи гвоздями не прибиты (только bucket)_

`local`: 
+ go install github.com/go-task/task/v3/cmd/task@latest
+ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
+ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
+ go install github.com/swaggo/swag/cmd/swag@latest
+ go install golang.org/x/tools/cmd/goimports@latest
+ go install mvdan.cc/gofumpt@latest
+ docker compose up --profile deps -d (поднимет s3, psql, redpanda)

//TODO: в task'ах написать форматеры для sql + proto + go + smt..?
//TODO: подумать над решение задачи об общей /pkg для всех сервисов
//TODO: сделать так, что бы панда показывала proto сообщения