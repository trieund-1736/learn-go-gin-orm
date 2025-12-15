# API Test Instructions

## List all users

```bash
curl -X GET http://localhost:8080/api/v1/users
```

## Create new user

```bash
curl -X POST http://localhost:8080/api/v1/users \
-H "Content-Type: application/json" \
-d '{"name": "John Doe", "email": "john.doe@example.com", "password": "password123"}'
```

## Get user by ID

```bash
curl -X GET http://localhost:8080/api/v1/users/1
```

## Update user

```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
-H "Content-Type: application/json" \
-d '{"name": "John Updated", "email": "john.updated@example.com"}'
```

## Delete user

```bash
curl -X DELETE http://localhost:8080/api/v1/users/1
```
