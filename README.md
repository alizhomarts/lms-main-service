# LMS Main Service

LMS Main Service is a backend service for managing courses, chapters, and lessons in an LMS system.

The service is built with **Go**, **Gin**, **GORM**, **PostgreSQL**, **Goose migrations**, **Swagger**, and **Docker**.

---

## Features

- Course CRUD
- Chapter CRUD
- Lesson CRUD
- PostgreSQL integration
- Goose database migrations
- Swagger API documentation
- Clean layered architecture
- Centralized error handling
- Unified JSON response format
- Docker Compose support

---

## Tech Stack

- Go
- Gin
- GORM
- PostgreSQL
- Goose
- Swagger / Swaggo
- Docker
- Docker Compose
- Logrus

---

## Project Structure

```text
lms-main-service/
│
├── cmd/
│   └── app/
│       └── main.go
│
├── internal/
│   ├── apperror/
│   ├── config/
│   ├── database/
│   ├── dto/
│   ├── entity/
│   ├── handler/
│   ├── middleware/
│   ├── repository/
│   ├── response/
│   ├── router/
│   └── service/
│
├── migrations/
├── docs/
├── .env.example
├── .gitignore
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md