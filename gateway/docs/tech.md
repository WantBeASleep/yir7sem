## Tech

### Что делает api-gateway?

1) Проверяет подпись jwt token'а - нужно для авторизации 
2) Использует id из jwt token для дальнейших запросов (на фронт открыта ручка /getDoctorPatients, id доктора подставит gateway из распашенного jwt)
3) Добавляет в ctx/метаданные всех запросов x-request_id, нужно для сквозного межсервисного логирования
4) Реализует логику загрузок image/uzi через ручку /download
5) Реализует бизнес логику завязанную на несколько вызовов

### Термины
__ID DOCTOR = ID USER__
x-user_id - id извлеченный из jwt

__CARDS - СВЯЗЬ МЕЖДУ ДОКТОРОМ И ПАЦИЕНТОМ, ПОНЯТИЯ ОБЩЕЙ КАРТЫ НЕТ!!!!!__

### Ручки

+ POST /auth/register
    - вызовет auth /register
    - вызовет med /registerDoctor __НУЖНА РАСПРЕДЕЛЕННАЯ ТРАНЗАКЦИЯ__
+ POST /auth/login
    - вызовет auth /login
+ POST /auth/refresh
    - вызовет auth /refresh


+ GET /med/doctors
    - распарсит jwt, сравнит с id запроса, //TODO: сделать систему ролей для админки
    - вызовет med /getDoctor

+ PATCH /med/doctors
    - распарсит jwt, сравнит с id запроса
    - вызовет med /updateDoctor

+ GET /med/doctors/patients
    - распарсит jwt, сравнит с id запроса
    - вызовет med /getDoctorPatients


+ POST /med/patients
    - вызовет med /createPatient

+ PATCH /med/patients/{id}
    - распарсит jwt, извлечет id для запроса (валидация возможности имзенения на стороне сервиса)
    - вызовет med /updatePatient

+ GET /med/patients/{id}
    - вызовет med /getPatient


+ POST /med/cards __ДОБАВИТ ПАЦИЕНТА К ТЕКУЩЕМУ ВРАЧУ__
    - распарсит jwt, извлечет id для запроса
    - вызовет med /createCard

+ PATCH /med/cards/{id}
    - распарсит jwt, извлечет id для запроса
    - вызовет med /updateCard

+ GET /med/cards/{id}
    - распарсит jwt, извлечет id для запроса
    - вызовет med /getCard


//TODO: вопрос, может ли врач, если это не его пациент загружать снимок?
//TODO: если нет, сделать валидацию
+ POST /uzi/uzis
    - распарсит jwt, извлечет id для запроса
    - вызовет uzi /createUzi
    - загружит в s3 узи
    - отправит ивент в брокер

+ PATCH /uzi/uzis/{id}
    - распарсит jwt, извлечет id для запроса
    - вызовет uzi /updateUzi

+ PATCH /uzi/echographics/{id}
    - распарсит jwt, извлечет id для запроса
    - вызовет uzi /updateEchographic

+ GET /uzi/uzis/{id}
    - распарсит jwt, извлечет id для запроса
    - вызовет uzi /getUzi

+ GET /uzi/uzis/{id}/images
    - распарсит jwt, извлечет id для запроса
    - вызовет uzi /getUziImages

+ GET /uzi/images/{id}/nodes-segments
    - распарсит jwt, извлечет id для запроса
    - вызовет uzi /getImageSegmentsWithNodes


+ GET /uzi/devices
    - распарсит jwt, извлечет id для запроса
    - вызовет uzi /getDeviceList



+ POST /uzi/nodes
    - распарсит jwt, извлечет id для запроса
    - вызовет uzi /createNode

+ DELETE /uzi/nodes/{id}
    - распарсит jwt, извлечет id для запроса
    - вызовет uzi /deleteNode

+ PATCH /uzi/nodes/{id}
    - распарсит jwt, извлечет id для запроса
    - вызовет uzi /updateNode



+ POST /uzi/segments
    - распарсит jwt, извлечет id для запроса
    - вызовет uzi /createSegment

+ DELETE /uzi/segments
    - распарсит jwt, извлечет id для запроса
    - вызовет uzi /deleteSegment

+ PATCH /uzi/segments
    - распарсит jwt, извлечет id для запроса
    - вызовет uzi /updateSegment


    

