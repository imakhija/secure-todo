# Secure Todo

A simple, secure full-stack todo application built with Go, Gin, PostgreSQL, and JWT authentication.

## Project Structure

```text
secure-todo/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── auth/
│   │   ├── jwt.go
│   ├── db/
│   │   ├── db.go
│   │   └── user_repository.go
│   ├── handlers/
│   │   └── auth.go
│   └── models/
│       ├── user.go
│       └── todo.go
├── migrations/
│   └── 001_init.sql
├── web/
├── docker-compose.yml
├── .env
├── go.mod
└── go.sum
```

---

## Local Development

### Clone this repository
`git clone https://github.com/imakhija/secure-todo.git`

### Install dependencies
`go mod tidy`

### Configure .env
```
POSTGRES_USER=...
POSTGRES_PASSWORD=...
POSTGRES_DB=...
DATABASE_URL=...
JWT_SECRET=...
```

### Start PostgreSQL

```bash
docker compose up -d
```

### Verify Database

```bash
docker ps
```

### Run Application

```bash
go run ./cmd/server
```

Application runs on:

```text
http://localhost:8080
```

Health endpoint:

```text
GET /health
```

---

## API

### Register User

```http
POST /api/register
```

Request:

```json
{
  "username": "imakhija",
  "password": "12345"
}
```

Successful Response:

```json
{
  "id": 1
}
```

### Login User

```http
POST /api/login
```

Request:

```json
{
  "username": "imakhija",
  "password": "12345"
}
```

Successful Response:

```json
{
  "token": "eyJ..."
}
```