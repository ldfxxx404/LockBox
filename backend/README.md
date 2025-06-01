# LockBox Backend Docs

## Config

### ENV variables

```bash
SERVER_PORT        Server listening port
JWT_SECRET         JWT secret key
POSTGRES_URL       PostgreSQL connection URL
MIGRATIONS_DIR     Path to database migration files
DB_NAME            Database name
STORAGE_LIMIT      User storage limit (in megabytes)
MINIO_ENDPOINT     MinIO server address
MINIO_ACCESS_KEY   MinIO access key
MINIO_SECRET_KEY   MinIO secret key
MINIO_BUCKET       MinIO bucket name
MINIO_USE_SSL      Enable SSL for MinIO (true/false)
```

### Fiber

```bash
FiberConfig        Config for body size limits and other options
```

### Limiter

```bash
Limiter            Rate limiting settings (requests per period)
```

---

## API Endpoints v1 (test)

You can find full API documentation at:

```
http://localhost:5000/docs/index.html
```

Swagger files are located in `cmd/swagger/`.

---

## Project Structure and Requirements

### Requirements

- [Go](https://go.dev/) >= 1.24
- [Fiber](https://github.com/gofiber/fiber) — HTTP framework
- [golang-jwt](https://github.com/golang-jwt/jwt) and [gofiber/jwt](https://github.com/gofiber/jwt) — authentication
- [swaggo/swag](https://github.com/swaggo/swag), [fiber-swagger](https://github.com/swaggo/fiber-swagger) — Swagger documentation
- [PostgreSQL](https://www.postgresql.org/), [sqlx](https://github.com/jmoiron/sqlx) — relational database
- [goose](https://github.com/pressly/goose) — SQL migrations
- [minio-go](https://github.com/minio/minio-go) — S3-compatible file storage
- [chambrracelet](https://github.com/charmbracelet/log) — logging
- [Docker](https://www.docker.com/) — container runtime
- [docker-compose](https://docs.docker.com/compose/) — service orchestration
- [make](https://www.gnu.org/software/make/) — build automation
- [git](https://git-scm.com/) — version control

### Modules

- `main` — entrypoint of the backend
- `swagger` — auto-generated Swagger docs
- `config` — environment and application configuration
- `database` — DB initialization and connection
- `handlers` — HTTP route logic
- `middleware` — authentication and request filtering
- `models` — DTOs and data models
- `repositories` — data access layer
- `services` — business logic (file manager, user services)
- `utils` — helper functions and internal utilities

### Miscellaneous

- `migrations/` — SQL migration files
- `Dockerfile` — Docker image build script
- `docker-compose.yml` — (if present) container setup for services

## Logs flags

- `--info` — info log level
- `--debug` — debug log level
- `--errlog` — error log level
