# Тестовое от "Effective Mobile"
# Реализация онлайн библиотеки песен 🎶

## Build
Создайте .env файл
```bash
cp .env.example .env
```

```bash
make build
```

## Dependencies
```bash
go mod tidy
```

## Start

```bash
bin/app
```


## Docker
```bash
docker-compose up --build -d
```

## Swagger link
http://${your_host}:${PORT}/swagger/index.html

# Environment

| env                         | default value          | description                                |
|:----------------------------|:-----------------------|:-------------------------------------------|
| PORT                        | 3000                   | Service port                               |
| APP_ENV                     | development            | App environment (development or production)|
| POSTGRES_HOST               | localhost              | Postgres host                              |
| POSTGRES_PORT               | 5432                   | Postgres port                              |
| POSTGRES_DB_NAME            | core                   | Postgres database name                     |
| POSTGRES_USER               | postgres               | Postgres user                              |
| POSTGRES_PWD                | root                   | Postgres password                          |
