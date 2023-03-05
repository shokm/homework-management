# README

## 使用方法

1. まず、登録ページでユーザー登録をします。
2. 課題一覧ページに遷移するので、課題を追加します。
3. 教科を追加したい場合は教科一覧ページに遷移し、教科を追加します。
4. 課題が完了したら、各課題の左端にあるボタンを押し、課題を完了させます。

## 技術

### Frontend

- Nuxt.js
    - @nuxtjs/pwa
    - @nuxtjs/auth-next

### Backend

- Golang
    - dgrijalva/jwt-go
    - go-gorm/gorm
- Swagger (go-swagger)
- Docker (PostgreSQL)

## 起動方法（開発時）

Node バージョン

```
node -v
v16.13.0
```

Go バージョン

```
go version
go version go1.18.3 darwin/arm64
```

Swagger バージョン

```
swagger version
version: v0.29.0
```

必要なパッケージをインストール

```
go mod tidy
```

Swaggerコード生成

```
swagger generate server -t gen -f swagger/swagger.yaml
```

docker composeでDB作成

```
cd settings
docker-compose build
```

DB Migrate

```
go run settings/database_migrate.go
```

backend 環境変数設定（backend/settings/.env）

- JWT_SECRET_KEY
    - 任意の秘密鍵を設定
- DATA_SOURCE_NAME
    - DBの接続情報を設定

```
JWT_SECRET_KEY="YOUR_SECRET_KEY"
DATA_SOURCE_NAME="host=localhost user=postgres password=postgrespw dbname=postgres port=55000 sslmode=disable TimeZone=Asia/Tokyo"
```

frontend 環境変数設定（frontend/.env）

- API_URL
    - backendのURLを入力

```
API_URL='http://example.com/'
```

backend起動

```
go run gen/cmd/homework-management-server/main.go --host 0.0.0.0 --port 8888
```

frontend起動

```
npm run dev
```