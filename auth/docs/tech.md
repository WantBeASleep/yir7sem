## Tech

### Сущности их отношеня и ручки

#### User
* id _соответствует с id мед работна_
* email 
* password _в захешированным виде_
* tokem _refresh token_

/register
    - -> email, password
    - <- id

/login
    - -> email, password
    - <- access_token, refresh_token

/refresh
    - -> refresh_token
    - <- access_token