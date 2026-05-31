# ディレクトリ構成

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
    ├── library/ // ※将来的に作成。実DBクライアント APIクライアントの外部依存のラッパー
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

## 構成理由

### 要件

- 小～中規模を想定
- 技術と業務が分割できること
- アプリケーションとinfraコードを同居させてAI駆動開発しやすくすること
- Lambda + API Gateway構成のアプリケーションをデプロイすることを念頭におく

### 理由

- ディレクトリ構成にはクリーンアーキテクチャを採用することで技術と業務を分割しやすくする
- `internal` 配下に機能群でディレクトリを置くことで将来的な拡張性を考慮
- 共通処理は `internal/common` 配下に置き、複数機能から利用する境界を明確にする。
- `cmd` 配下に実行エンドポイントを配置することでアプリケーションとエンドポイントを分離。
- `build` 配下に Dockerfile を配置することでアプリケーションコードとビルド定義を分離し、機能・デプロイ単位を揃える。
- `internal/presentation/<機能>` は API の入口を置き、将来複数機能の `application` をまたいで扱えるようにする。
- `internal/library` に外部依存のラッパーを置くことでテストダブルが作成しやすいようにする。

### 対象外としたもの

- DIは本筋から外れるため今回は簡潔さを優先して扱わない
- PoCなので実DBは本リポジトリでは扱わず、ベタ書きとした。
