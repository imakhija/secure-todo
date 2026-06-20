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
│   │   └── jwt.go
│   ├── db/
│   │   ├── db.go
│   │   └── user_repository.go
│   ├── handlers/
│   │   └── auth.go
│   ├── middleware/
│   │   └── auth.go
│   └── models/
│       ├── user.go
│       └── todo.go
├── migrations/
│   └── 001_init.sql
├── web/
│   └── index.html
├── docker-compose.yml
├── .env
├── go.mod
└── go.sum
```

---

## Local Development

### Clone this repository

```bash
git clone https://github.com/imakhija/secure-todo.git
cd secure-todo
```

### Install dependencies

```bash
go mod tidy
```

### Configure `.env`

Create a `.env` file in the project root:

```env
POSTGRES_USER=secure_todo
POSTGRES_PASSWORD=secure_todo
POSTGRES_DB=secure_todo
DATABASE_URL=postgres://secure_todo:secure_todo@localhost:5432/secure_todo?sslmode=disable
JWT_SECRET=change-me-to-a-long-random-string
```

`JWT_SECRET` is required for login and for validating tokens on protected routes. Tokens expire after 24 hours.

### Start PostgreSQL

```bash
docker compose up -d
```

### Apply migrations

Run the initial schema against the database (adjust `-U` and `-d` to match your `.env`):

```bash
cat migrations/001_init.sql | docker compose exec -T postgres psql -U secure_todo -d secure_todo
```

### Run the application

```bash
go run ./cmd/server
```

The server listens on:

```text
http://localhost:8080
```

Health check:

```http
GET /health
```

Response:

```json
{
  "status": "healthy"
}
```

---

## Authentication

Login returns a JWT signed with `JWT_SECRET`. Protected routes under `/api` require this token in the `Authorization` header:

```http
Authorization: Bearer eyJ...
```

The auth middleware validates the token, checks expiry, and attaches the authenticated `user_id` to the request context. Invalid or missing tokens receive `401 Unauthorized`.

---

## API

### Register User

```http
POST /api/register
Content-Type: application/json
```

Request:

```json
{
  "username": "imakhija",
  "password": "12345"
}
```

Successful response (`201 Created`):

```json
{
  "id": 1
}
```

### Login User

```http
POST /api/login
Content-Type: application/json
```

Request:

```json
{
  "username": "imakhija",
  "password": "12345"
}
```

Successful response (`200 OK`):

```json
{
  "token": "eyJ..."
}
```