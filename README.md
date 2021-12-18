# Go PostgreSQL

---

Этот репозиторий представляет собой сервис, который по REST API получает запросы на добавление карточек продуктов и запросы на их поиск

## Используемы технологии
- Go
- Elasticsearch

## Запуск проекта
- склонировать репозиторий
- выполнить `docker-compose up -d`
- api доступно на `http://localhost:3000/`


## Документация по методам

- [Продукты](#Продукты)
    - [Добавление продукта](#Добавление-продуктов)
    - [Поиск по продуктам](#Поиск-продуктов)

## Продукты
### Добавление продуктов

`POST /api/v1/product` - добавление нового продукта.

__Обязательные параметры запроса:__

- `title` - название продукта
- `description` - описание продукта
- `price` - цена продукта
- `quantity` - количество

__Пример запроса:__

    curl -X POST http://localhost:3000/api/v1/product \
    -H 'Content-Type: application/json' \
    -d '{"title": "Электрическая зубная щетка", "description": "Звуковая технология в сочетании с конструкцией нашей зубной щетки обеспечивает в 3 раза более эффективное удаление налета, по сравнению с обычной зубной щеткой.", "price": 4700, "quantity": 10 }'


### Поиск продуктов

`GET /api/v1/product/search` - поиск товаров в БД

__Обязательные параметры запроса:__

- `q` - строка запроса

__Пример запроса:__

    curl -X GET http://localhost:3000/api/v1/product/search?q=щет


## Road map

- Тесты