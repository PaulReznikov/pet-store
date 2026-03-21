# Pet Store Backend API

Backend приложение для зоомагазина, реализующее REST API для управления товарами и операциями магазина.

## Описание

Pet Store - это backend система для управления зоомагазином, предоставляющая полный набор функций для работы с каталогом товаров, включая создание, чтение, обновление и удаление записей (CRUD операции), а также импорт и экспорт данных о товарах.

## Функциональность

- ✅ **CRUD операции для товаров**
  - Создание новых товаров
  - Получение списка товаров с фильтрацией и пагинацией
  - Получение информации о конкретном товаре
  - Обновление данных товара
  - Удаление товаров

- 📦 **Импорт/Экспорт данных**
  - Массовая загрузка товаров из JSON файлов
  - Выгрузка каталога товаров в различных форматах
  - Валидация данных при импорте

- 🔍 **Дополнительные возможности**
  - Поиск товаров по различным параметрам
  - Фильтрация по категориям (корм, игрушки, аксессуары и т.д.)
  - Фильтрация по брендам
  - Сортировка товаров:
    - По дате добавления
    - По названию: от А до Я
    - По названию: от Я до А
    - По цене по возрастанию
    - По цене по убыванию
    - По популярности
  - Управление остатками на складе
  - История изменений товаров

## Технологический стек

- **Язык:** Go 1.24 или выше
- **База данных:** PostgreSQL
- **API документация:** Swagger/OpenAPI
- **Миграции:** goose
- **Конфигурация:** ENV файлы / YAML

## Архитектура проекта

```
pet-store/
├── cmd/                    # Точки входа приложения
│   └── main.go            # Главный файл запуска
├── internal/              # Внутренняя логика приложения
│   ├── api/               # API слой
│   │   ├── interfaces/    # Интерфейсы и контракты
│   │   ├── models/        # DTO модели
│   │   ├── services/      # Бизнес-логика
│   │   └── transport/     # HTTP handlers
│   ├── app/               # Инициализация приложения
│   ├── config/            # Конфигурация
│   ├── core/              # Основная бизнес-логика
│   └── infrastructure/    # Работа с БД, внешними сервисами
├── pkg/                   # Публичные библиотеки
├── migrations/            # Миграции базы данных
├── tests/                 # Тесты
├── deploy/                # Конфигурации для развертывания
└── local/                 # Локальное окружение
```

## Требования

- Go 1.24 или выше
- PostgreSQL 14+
- Docker и Docker Compose (опционально)

## Установка и запуск

### Локальная разработка

1. Клонируйте репозиторий:
```bash
git clone https://github.com/PaulReznikov/pet-store.git
cd pet-store
```

2. Установите зависимости:
```bash
go mod download
```

3. Настройте конфигурацию:
```bash
cp .env.example .env
# Отредактируйте .env файл с вашими настройками
```

4. Примените миграции базы данных:
```bash
make migrate-up
# или
cd migrations && goose up
```

5. Запустите приложение:
```bash
make run
# или
go run cmd/main.go
```

### Запуск через Docker (пока в разработке)

```bash
docker-compose up -d
```

## API Endpoints (пока в разработке)

### Товары (Products)

| Метод | Endpoint | Описание |
|-------|----------|----------|
| GET | `/api/v1/popular_products` | Получить список популярных товаров |
| GET | `/api/v1/products/:id` | Получить товар по ID |
| POST | `/api/v1/products` | Создать новый товар |
| PUT | `/api/v1/products/:id` | Обновить товар |
| DELETE | `/api/v1/products/:id` | Удалить товар |



## Конфигурация

Пример `.env` файла:

```env
# Server
SERVER_HOST=0.0.0.0
SERVER_PORT=8080

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=petstore
DB_PASSWORD=your_password
DB_NAME=petstore_db
DB_SSLMODE=disable

# Goose (для миграций)
GOOSE_DRIVER=postgres
GOOSE_DBSTRING=host=localhost port=5432 user=petstore password=your_password dbname=petstore_db sslmode=disable
GOOSE_MIGRATION_DIR=./migrations

# Application
APP_ENV=development
LOG_LEVEL=debug

# File Upload
MAX_UPLOAD_SIZE=10485760  # 10MB
ALLOWED_FORMATS=csv,json
```

## Тестирование

Запуск всех тестов:
```bash
make test
```

Запуск тестов с покрытием:
```bash
make test-coverage
```

Запуск линтера:
```bash
make lint
```


## Swagger документация

После запуска приложения документация API доступна по адресу:
```
http://localhost:8080/swagger/
```

## Разработка

### Работа с миграциями (goose)


Создание новой миграции (с timestamp):
```bash
cd migrations
goose create add_new_table sql
```

Применение миграций:
```bash
goose up
```

Откат последней миграции:
```bash
goose down
```

Проверка статуса миграций:
```bash
goose status
```

## Безопасность

- Валидация всех входных данных
- Защита от SQL инъекций
- Rate limiting для API endpoints
- Логирование всех операций
- (Опционально) JWT аутентификация


## Roadmap

- [ ] Аутентификация и авторизация
- [ ] Управление заказами
- [ ] Интеграция с платежными системами
- [ ] Управление поставщиками
- [ ] Система скидок и акций
- [ ] Уведомления о низких остатках
- [ ] REST API для мобильного приложения

## Контакты и поддержка
- Авторы: Paul Reznikov [GitHub](https://github.com/PaulReznikov), Artsem Yunchyts [GitHub](https://github.com/Oliver1ck)
- Issues: [GitHub Issues](https://github.com/PaulReznikov/pet-store/issues)

## Лицензия

MIT License

---

**Примечание:** Данный проект находится в разработке. Некоторые функции могут быть не полностью реализованы.
