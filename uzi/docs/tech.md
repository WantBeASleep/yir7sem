## Tech

### Сущности ручки

! _Пути для ручек - именование gRPC методов_

#### Device

Структура:
* id
* name

Обозначения:
    - device: id, name

+ /getDeviceList 
    - <- []device

+ /createDevice _private_ _testing_
    - -> name
    - <- id

#### Uzi

Структура:
* id
* projection
* patient_id
* device_id

Обозначения:
    - Uzi: id, projection, patient_id, device_id
    - CreateUziReq: projection, patient_id, device_id
    - UziMut: projection, patient_id

+ /createUzi
    - -> CreateUziReq
    - <- id

+ /updateUzi 
    - -> id + UziMut
    - <- Uzi

+ /getUzi
    - -> id
    - <- Uzi

#### Image

Структура:
* id
* page
* uzi_id

Обозначения:
    - image: id, page

+ /getUziImages
    - -> uzi_id
    - <- []image

+ /getNodesWithSegmentsOnImage
    - -> id
    - <- []node{id, ai, tirads_23, tirads_4, tirads_5}, []segments{id, node_id, image_id, contor, tirads_23, tirads_4, tirads_5}

__Kafka__

+ uziUploaded
    - -> uzi_id
        * Выгрузить из S3
        * Split картинки
        * Загрузка каждой в S3
        * Сохранение в бд
        * Написание в event uziSplitted
            - -> uzi_id, []image_id

#### Segment
Структура:
* id
* node_id
* image_id
* contor
* tirads_23
* tirads_4
* tirads_5

Обозначения:
    - Segment: id, node_id, image_id, contor, tirads_23, tirads_4, tirads_5
    - AddSegmentReq: node_id, image_id, contor, tirads_23, tirads_4, tirads_5
    - SegmentMut: tirads_23, tirads_4, tirads_5

+ /addSegment
    - -> AddSegmentReq
    - <- id

+ /delSegment
    - -> id

+ /updateSegment
    - -> id, SegmentMut
    - <- Segment


#### Node
Структура:
* id
* ai
* tirads_23
* tirads_4
* tirads_5

Обозначения:
    - Node: id, ai, tirads_23, tirads_4, tirads_5
    - CreateNode: id, ai, tirads_23, tirads_4, tirads_5, []NestedToNodeSegment: image_id, contor, tirads_23, tirads_4, tirads_5
    - NodeMut: tirads_23, tirads_4, tirads_5

+ /createNode
    - -> CreateNode
    - <- id

+ /delNode
    - -> id

+ /updateNode
    - -> id, NodeMut
    - <- Node

#### Echographics

Является просто эхопризнаками узи с отношением `1:1`
есть ручка update + получается из get'а uzi 

## TODO:
+ Keep alive