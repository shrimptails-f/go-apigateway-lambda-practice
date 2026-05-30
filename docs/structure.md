# Internal Structure

```text
root/
├── build/
│   └── user/
│       └── Dockerfile
├── infra/
└── internal/
    └── user/
        ├── application/
        │   ├── get_detail_usecase.go
        │   └── list_usecase.go
        ├── domain/
        │   └── user.go
        ├── infrastructure/
        │   └── repository/
        │       └── in_memory.go // PoCなのでDB接続は行わずベタ書きで返す
        ├── presentation/
        │   ├── handler.go
        │   ├── response.go
        │   └── router.go
        └── main.go
```
