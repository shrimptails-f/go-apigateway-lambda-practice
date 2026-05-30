# Internal Structure

```text
root/
├── cmd/
│   └── api/
│       └── user/
│           └── main.go // 機能・デプロイ単位で分ける
├── build/
│   └── user/
│       └── Dockerfile // 機能・デプロイ単位で分ける
├── infra/
└── internal/
    ├── common/
    │   └── presentation/
    │       └── response.go
    ├── presentation/
    │   └── user/
    │       ├── application_interface.go
    │       ├── handler.go
    │       ├── response.go
    │       └── router.go
    └── user/
        ├── application/
        │   ├── get_detail_usecase.go
        │   └── list_usecase.go
        ├── domain/
        │   ├── error.go
        │   └── user.go
        └── infrastructure/
            └── repository/
                └── in_memory.go // PoCなのでDB接続は行わずベタ書きで返す
```
