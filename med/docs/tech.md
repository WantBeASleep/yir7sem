## Tech
### Сущности, их отношения, ручки

#### Doctor

__Doctor__: это сущность описывающая самого врача, а не его аккаунт. По сути это его мед принадлежность Org && Job.

Структура:
* id - Совпадает с id _user из auth service_
* FullName _ФИО_
* Org _Учесть вынос в отдельную табличку_
* Job _Должность_
* Desc _Опыт работы + описание_

обозначения
    doctor: id, fullname, org, job, desc

+ /registerDoctor
    - -> doctor
Регистрация будет осуществоляться после реги в __Auth service__, значит потребуется распределенная транзакция для регистрации.

+ /getDoctor
    - -> id
    - <- doctor

_patch_
+ /updateDoctor
    - -> id, org, job, desc
    - <- doctor

#### Patient

__Patient__: сущность описывающая пациента как человека, его личные данные и характеристики
Сюда относится _общая_ для всех врачей информация, обн/не обн образование/активный пациент.
Все узи будут закрепляться именно на пациенте

//TODO: надо синкнутьсяч на счет имя лечащего врача

Структура:
* id 
* fullname   
* email  
* policy _мед полис_
* active _автивен ли пациент_
* malignancy _образование_
* last_uzi_date _дата последнего узи снимка_ //TODO: связать логику с uzi

обозначения:
    patient: id, fullname, email, policy, active, malignancy, last_uzi_date

__РУЧКА НЕ ПРИВЯЗЫВАЕТ ПАЦИЕНТА К ДОКТУРУ, ДЛЯ ЭТОГО ЕСТЬ CARD__
+ /createPatient
    - -> fullname, email, policy, active, malignancy
    - <- patient

+ /updatePatient
    - -> id, doctor_id, active, malignancy
    - <- patient
__РУЧКА ПРОВЕРИТ ЧТО doctor имеет доступ к этому (его пациент)__

+ /getPatient
    - -> id
    - <- patient

//TODO: перенести ручку в doctor
+ /getDoctorPatients
    - -> doctor_id
    - <- []patient

__внес в /updatePatient__
+ /updatePatientLastUzi _ручка обновит последние время узи у пациента и выкинет ошибку, если более актуальная версия уже есть
    - -> id, last_uzi_date


__СИЛЬНО ОТЛИЧАЕТСЯ ОТ СТАРЫХ РУЧЕК__
Сначала создаем пациента, и потом на него подвязываем врачей.

Альтернатива при добавления снимка проверять полис на уникальность, но теже яйца только с скрытой логикой - не вижу смысла в этом

#### Card

Сущность __Card__ отвечает за взаимодействие __КОНКРЕТНОГО__ врача с __КОНКРЕТНЫМ__ пациентом. По сути реализует связь `M:N`, но с добавлением инфы в нее

//TODO: добавить по все таблицы create_at, update_at

//TODO: пробежаться и поправить имена в стурктурах
Структура:
* doctor_id
* patient_id
* diagnosis

+ /createCard
    - -> doctor_id, patient_id, diagnosis

//TODO: добавить ошибки для случаев когда не найдено ВЕЗДЕ!
+ /updateCard
    - -> doctor_id, patient_id, diagnosis
    - <- doctor_id, patient_id, diagnosis

+ /getCard
    - -> doctor_id, patient_id
    - <- doctor_id, patient_id, diagnosis
