## Структура
```
auth
├── api <-- все что касается походов в сервис из вне
│  └── auth <-- proto auth сервиса
│     └── auth.proto
├── cmd <-- вход в приложение
│  ├── auth
│  │  └── main.go
│  └── tools
│     └── main.go
├── config <-- конфиг
│  └── config.yml
├── internal <-- папка с внутренним кодом приложения, не импортируемым
│  ├── apps <-- контроллер, поднимающий порты и контроллеры приложения
│  │  └── auth.go
│  ├── config <-- структура конфига
│  │  └── config.go
│  ├── controller <-- контроллер слой отвечающий за взаимодействие с приложением
│  │  ├── auth <-- контроллер auth
│  │  │  └── auth.go
│  │  ├── usecases <-- интерфейсы для взаимодействием с юзкейсами
│  │  │  └── auth.go
│  │  └── validation <-- валидация поступающий данных
│  │     ├── login.go
│  │     ├── refresh.go
│  │     ├── register.go
│  │     └── validator.go
│  ├── core <-- кор слой, корневая бизнес логика
│  │  └── jwt <-- jwt сервис, выдает и проверяет токены
│  │     └── jwt.go
│  ├── entity <-- entity слой, хранит структуры для общения меж слоями/алгоритмы
│  │  ├── entity.go
│  │  ├── errors.go
│  │  └── hash.go
│  ├── repositories <-- repo слой, для взаимодействия с внешними системами
│  │  ├── db <-- взаимодействия с базами данных в repo слое
│  │  │  ├── mappers <-- мапперы перевода entity в repo структуры
│  │  │  │  └── universal_mappers.go
│  │  │  ├── models <-- модели базы данных
│  │  │  │  └── auth.go
│  │  │  ├── repositories <-- код взаимодействия с бд
│  │  │  │  └── auth.go
│  │  │  └── utils <-- утилити для бд
│  │  │     └── dns.go
│  │  └── services <-- взаимодействия с внешними сервисами
│  │     └── med.go
│  └── usecases <-- слой верхней логики "оркестрации" 
│     ├── auth <-- юзкейсы аутентификации
│     │  ├── auth.go
│     │  └── tokens.go
│     ├── core <-- интерфейсы взаимодействия с core слоем
│     │  └── jwt.go
│     └── repositories <-- интерфейсы взаимодействия с репо слейм
│        ├── auth.go
│        └── med.go
├── README.md
└── Taskfile.yml <-- файл сборки
```
