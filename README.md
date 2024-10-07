# garik-api

## Build

Создайте .env файл
```bash
cp .env.example .env
```
Убедитесь, что верно заполнили переменные окружения

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


# Environment

| env                         | default value          | description                                |
|:----------------------------|:-----------------------|:-------------------------------------------|
| PORT                        | 8080                   | Service port                               |
| APP_ENV                     | development            | Type of instance(development | production) | 
| OPENAI_COMPLETION_MODEL     | gpt-4o                 | Open AI credentials                        | 
| OPENAI_API_KEY              |                        | Open AI credentials                        | 
| OPENAI_URL                  | https://api.openai.com | Open AI credentials                        | 
