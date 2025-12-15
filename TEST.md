# API Test Instructions

## List all users

```bash
curl -X GET http://localhost:8080/api/v1/users | jq .
```

## Create new user

```bash
curl -X POST http://localhost:8080/api/v1/users \
-H "Content-Type: application/json" \
-d '{"name": "John Doe", "email": "nguyen.van.a@sun-asterisk.com"}' | jq .
```

## Get user by ID

```bash
curl -X GET http://localhost:8080/api/v1/users/1 | jq .
```

## Update user

```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
-H "Content-Type: application/json" \
-d '{"name": "John Updated", "email": "nguyen.van.a.update@sun-asterisk.com"}' | jq .
```

## Delete user

```bash
curl -X DELETE http://localhost:8080/api/v1/users/1 | jq .
```
