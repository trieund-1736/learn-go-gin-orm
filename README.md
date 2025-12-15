# learn-go-gin-orm

A REST API application for managing users built with **Go**, **Gin framework**, and **GORM ORM**.

## Objective

Practice building REST APIs with Gin and GORM, implementing complete CRUD operations with proper error handling and pagination.

## Features

- ✅ CRUD operations for users
- ✅ Pagination support
- ✅ Soft delete (using DeletedAt)
- ✅ JSON request/response handling
- ✅ Error handling with appropriate HTTP status codes
- ✅ Database migration with GORM

## Project Structure

```text
internal/
├── config/          # Database configuration
├── models/          # Data models (User)
├── repository/      # DB access layer
├── services/        # Business logic
├── handlers/        # HTTP handlers
└── routes/          # Route definitions
```

## Setup

### 1. Create MySQL Database

```bash
docker run -d --name learn-go-gin-orm \
  -p 3306:3306 \
  -e MYSQL_ROOT_PASSWORD=password \
  -e MYSQL_DATABASE=learn_go_gin_orm \
  mysql:latest
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Run Application

```bash
go run cmd/app/main.go
```

Server runs on `http://localhost:8080`

## Database

### Connection Details

- **Host**: localhost
- **Port**: 3306
- **User**: root
- **Password**: password
- **Database**: learn_go_gin_orm

### Users Table Schema

| Column     | Type         | Constraint                  |
| ---------- | ------------ | --------------------------- |
| id         | BIGINT       | PRIMARY KEY, AUTO_INCREMENT |
| name       | VARCHAR(255) | NOT NULL                    |
| email      | VARCHAR(255) | UNIQUE, NOT NULL            |
| phone      | VARCHAR(255) | NULL (optional)             |
| created_at | DATETIME     | AUTO_CREATE_TIME            |
| updated_at | DATETIME     | AUTO_UPDATE_TIME            |
| deleted_at | DATETIME     | NULL (soft delete)          |

## API Endpoints

| Method | Endpoint            | Description                      |
| ------ | ------------------- | -------------------------------- |
| GET    | `/api/v1/users`     | List all users (with pagination) |
| POST   | `/api/v1/users`     | Create new user                  |
| GET    | `/api/v1/users/:id` | Get user by ID                   |
| PUT    | `/api/v1/users/:id` | Update user                      |
| DELETE | `/api/v1/users/:id` | Delete user                      |

## Testing

See [TEST.md](TEST.md) for example curl commands to test all APIs.
