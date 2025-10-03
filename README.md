# WB Shortener

Сервис для сокращения ссылок с аналитикой переходов, кэшированием через Redis и веб-интерфейсом.

## Функционал

Сервис поддерживает:
- **POST /shorten** — создание новой сокращённой ссылки
- **GET /s/{short_url}** — переход по короткой ссылке
- **GET /analytics/{short_url}** — получение аналитики (число переходов, User-Agent, время переходов).

### Основные возможности
- Автоматическая генерация коротких ссылок
- Поддержка кастомных ссылок 
- Сбор аналитики: количество переходов, User-Agent, IP-адреса, Referer
- Агрегация статистики по дням и браузерам
- Кэширование ссылок через Redis для быстрого доступа
- Веб-интерфейс для удобного тестирования

## Установка и запуск проекта

### 1. Клонирование репозитория
```bash
git clone https://github.com/kstsm/wb-shortener
```

### 2. Настройка переменных окружения
Создайте `.env` файл, скопировав в него значения из `env.example`:
```bash
cp env.example .env
```

### 3. Запуск зависимостей (Docker)
```bash
make docker-up
```

### 4. Миграция базы данных
```bash
make goose-up
```

### 5. Запуск сервиса
```bash
make run
```

Сервис будет доступен по адресу: http://localhost:8080

## API запросов

### Создание короткой ссылки
```bash
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://example.com",
    "custom_alias": "my-link"
  }'
```

Ответ:
```json
{
  "result": {
    "short_url": "my-link",
    "original_url": "https://example.com"
  }
}
```

### Переход по короткой ссылке
```bash
curl -L http://localhost:8080/s/my-link
```

### Получение аналитики
```bash
curl http://localhost:8080/analytics/my-link
```

Ответ:
```json
{
  "result": {
    "short_url": "my-link",
    "original_url": "https://example.com",
    "total_clicks": 15,
    "daily_stats": [
      {
        "date": "2025-10-02",
        "clicks": 10
      },
      {
        "date": "2025-10-01",
        "clicks": 5
      }
    ],
    "monthly_stats": [
      {
        "month": "2025-10",
        "clicks": 15
      }
    ],
    "user_agent_stats": [
      {
        "user_agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
        "clicks": 8
      },
      {
        "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36",
        "clicks": 7
      }
    ]
  }
}
```

