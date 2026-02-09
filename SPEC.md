# CHLR Technical Specification (v0.1.0 MVP)

## 1. Core Architecture
Утилита представляет собой Single Binary CLI приложение, написанное на Go.
*   **CLI Framework:** `spf13/cobra`
*   **Template Engine:** `text/template` + `embed` (VFS)
*   **Execution:** Прямое взаимодействие с OS (FS, exec, signals)

## 2. Interface Contract

### Command: `init`
Инициализирует новый проект.

**Signature:**
`chlr init <module_name> [flags]`

**Arguments:**
*   `<module_name>` (string, required): Имя проекта (имя папки) и имя Go-модуля.

**Flags:**
*   `--mono` (bool, default: `false`):
    *   `false`: Standalone структура (все в корне).
    *   `true`: Структура монорепозитория (`backend/`, `frontend/`, `docker-compose` в корне).
*   `--db` (string, default: `none`):
    *   `none`: Без базы данных.
    *   `postgres`: Добавляет сервис PostgreSQL в docker-compose и драйвер `pgx` в Go.

## 3. Internal Logic Flow

1.  **Validation:** Проверка наличия аргумента `<module_name>`.
2.  **Environment Check:** Детекция версии Go на хосте (`runtime.Version()`).
3.  **FS Operations:** Создание директорий в зависимости от флага `--mono`.
4.  **Template Rendering:**
    *   Чтение шаблонов из `embed.FS`.
    *   Инъекция переменных (`ProjectName`, `GoVersion`, `UseDB`).
    *   Генерация файлов (`main.go`, `Dockerfile`, `docker-compose.yml`, `go.mod`).
5.  **Post-Processing:**
    *   Выполнение `go mod tidy` внутри созданного проекта для загрузки зависимостей (в т.ч. `github.com/Chelaran/yagalog`).

## 4. Generated Stack (The "Chelaran Standard")

*   **Go:** Standard Project Layout (`cmd/`, `internal/`).
*   **Router:** `net/http` (Go 1.22+ mux).
*   **Logger:** `github.com/Chelaran/yagalog`.
*   **Lifecycle:** Реализован паттерн Graceful Shutdown.
*   **Containerization:** Multi-stage Dockerfile (Alpine based).

## 5. Template Strategy

Директория `templates/` внутри бинарника:
*   `go-base/`: Базовый код (main.go, go.mod).
*   `docker/`: Dockerfile, docker-compose шаблоны с условиями `{{ if .UseDB }}`.
*   `config/`: .env, .gitignore.

## 6. Roadmap (Future)
*   [ ] Support for Python (FastAPI) generation.
*   [ ] Command `add <service>` (Redis, Kafka).
*   [ ] Persistent config `.chlr.yaml` for project state management.
