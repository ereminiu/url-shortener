# link-shorter
(scroll down for english description)

Аналог [Bitly](https://bitly.com)

## Функции:
* Получить короткий код для заданной ссылки (POST /addlink)
* Перейти по оригинальной ссылке по короткому коду (GET /getlink)
* Ввести кастомный код вместо случайного  (POST /addcustomlink)
* Получить исходную ссылку по кастомному коду (GET /getcustomlink)

## Инструкция по запуску: 
* Создайте базу данных в докер контейнере:
  ``` 
  docker run --name=shrt-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres
  ```
  Порты для базы данных можно задать в файле configs/config
* Запустите этот контейнер:
  ``` 
  docker exec -it id /bin/bash
  ``` 
* Если запускаете приложение в первый раз, то примените миграции:
  ``` 
  migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up
  ```
* Запустите main.go:
  ```
  go run cmd/main.go
  ```

## Описание:
* ### /addlink

формат запроса: 
```json
{
    "link": "https://github.com/avito-tech/auto-backend-trainee-assignment"
}
```
формат ответа: 
```json
{
    "hash": "0Mf61CHL"
}
```

* ### /getlink

формат запроса: 
```json 
{
    "code": "0Mf61CHL"
}
```

в ответ происходит redirect на исходную ссылку, если "code" не существует, возвращается ответ с кодом 406: 
```json
{
    "messange": "Hashcode not found"
}
```

* ### /addcustomlink

формат запроса: 
```json
{
    "link": "https://github.com/avito-tech/auto-backend-trainee-assignment",
    "custom_code": "avito-assigment"
}
```

Успешный ответ: 
```json
{
    "messange": "Your custom link created successfuly"
}
```
Если заданный код уже существует, то вернется ответ с кодом 406: 
```json
{
    "messange": "This custom link is already exist"
}
```
Если заданный код длиннее 40 символов, то вернется ответ с кодом 406:
```json
{
    "messange": "Failed to create your custom links: Your code should not exceed 40 characters"
}
```

* ### /getcustomlink
```json 
{
    "custom_code": "avito-assigment"
}
```
