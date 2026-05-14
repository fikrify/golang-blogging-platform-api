# Golang Blogging Platform API

Go implementation of the [Blogging Platform API](https://roadmap.sh/projects/blogging-platform-api) project from roadmap.sh.

REST API for a blogging platform built with Go, Gin, and SQLite.

## Tech Stack

- **Go** 1.26
- **Gin** — HTTP framework
- **GORM** — ORM
- **SQLite** — database

## Getting Started

### Prerequisites

- Go 1.26+

### Run

```bash
go run cmd/main.go
```

Server starts on `http://localhost:8080`.

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/posts` | Get all posts (supports `?term=` search) |
| POST | `/posts` | Create new post |
| GET | `/posts/:id` | Get post by ID |
| PUT | `/posts/:id` | Update post |
| DELETE | `/posts/:id` | Delete post |

## Project Structure

```
cmd/          # Entry point
internal/
  handler/    # HTTP handlers
  model/      # Data models
  repository/ # Database layer
  routes/     # Route definitions
  service/    # Business logic
database/     # DB setup
docs/         # OpenAPI spec
integration/  # Integration tests
```

## API Spec

OpenAPI spec available at `docs/blogging-platform-api.yaml`.