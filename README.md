# CHLR (Chelaran CLI)

**CHLR** — инструмент скаффолдинга и автоматизации разработки Full-Cycle проектов. Генерирует production-ready архитектуру, соответствующую инженерным стандартам агентства **Chelaran**.

## ⚡ Особенности
*   **Standard Layout:** Генерация структуры проекта согласно Go Standard Layout.
*   **Docker-First:** Автоматическое создание оптимизированных `Dockerfile` и `docker-compose.yml`.
*   **Zero-Dependency Router:** Использование нативного роутинга Go 1.22+.
*   **Chelaran Ecosystem:** Интеграция логгера [yagalog](https://github.com/Chelaran/yagalog) из коробки.
*   **Auto-Configuration:** Автодетект локальной версии Go для синхронизации окружений.

## 🛠 Установка

```bash
go install github.com/bambutcha/chlr@latest
```

## 🚀 Использование

### 1. Создание микросервиса (Standalone)
Создает Go-сервис с REST API и Graceful Shutdown.
```bash
chlr init github.com/user/my-service
```

### 2. Создание проекта с базой данных
Добавляет PostgreSQL в `docker-compose` и драйвер `pgx` в зависимости.
```bash
chlr init my-app --db=postgres
```

### 3. Режим монорепозитория
Создает структуру для раздельного бэкенда и единой инфраструктуры.
```bash
chlr init my-platform --mono --db=postgres
```

## 📂 Генерируемая архитектура (Standalone)

```text
my-app/
├── cmd/
│   └── api/
│       └── main.go        # Entry point + Graceful Shutdown
├── deployments/
│   └── Dockerfile         # Multi-stage build (Alpine)
├── .env                   # Environment variables
├── docker-compose.yml     # Infrastructure orchestration
├── go.mod                 # Module definition
└── Makefile               # Shortcuts
```

## ⚙️ Конфигурация (Flags)

| Флаг | Тип | Описание | Default |
| :--- | :--- | :--- | :--- |
| `--mono` | `bool` | Включить структуру монорепозитория (`backend/`) | `false` |
| `--db` | `string` | Подключить базу данных (`postgres`, `none`) | `none` |

---
**License:** MIT  
**Author:** Daniil Yagolnik (bambutcha)
