# Folder Structure

```text
my-go-project/
│
├── cmd/
│   └── app/
│       └── main.go              # Entry point (khởi động server, router, config)
│
├── go.mod                       # Go modules (dependencies)
├── go.sum
├── README.md
│
├── internal/                    # Code CHỈ dùng cho project này
│   ├── config/                  # Load & parse config
│   │   ├── config.go
│   │   └── database.go
│   │
│   ├── handlers/                # HTTP handlers / controllers
│   │   ├── user_handler.go
│   │   └── product_handler.go
│   │
│   ├── middleware/              # Middleware (auth, logging, recover...)
│   │   ├── auth.go
│   │   └── logger.go
│   │
│   ├── repository/              # DB access layer (CRUD, query)
│   │   ├── user_repository.go
│   │   └── product_repository.go
│   │
│   ├── models/                  # Data models / entities
│   │   ├── user.go
│   │   └── product.go
│   │
│   ├── services/                # Business logic (optional nhưng rất tốt)
│   │   ├── user_service.go
│   │   └── product_service.go
│   │
│   └── routes/                  # Route definitions
│       ├── user_routes.go
│       └── product_routes.go
│
├── pkg/                         # Package CÓ THỂ tái sử dụng cho project khác
│   ├── logger/
│   │   └── logger.go
│   └── validator/
│       └── validator.go
│
├── config/                      # File cấu hình
│   ├── app.yaml
│   └── database.yaml
│
├── env/                         # Config theo môi trường
│   ├── dev.env
│   ├── staging.env
│   └── prod.env
│
├── utils/                       # Hàm tiện ích chung
│   ├── time.go
│   └── string.go
│
├── helpers/                     # Helper cho workflow cụ thể
│   └── response_helper.go
│
├── static/                      # Static assets
│   ├── css/
│   ├── js/
│   └── images/
│
├── templates/                   # HTML templates
│   ├── layout.html
│   └── index.html
│
└── tests/                       # Unit / integration tests
    └── user_test.go

```
