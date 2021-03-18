# go-gin-test

Go 言語を用いた REST API の構築を試したリポジトリです。

## 使用ライブラリ

- [gin-gonic/gin](https://github.com/gin-gonic/gin#users)
- [go-gorm/gorm](https://github.com/go-gorm/gorm)
- [rubenv/sql-migrate](https://github.com/rubenv/sql-migrate)

## 実行

```
$ go run main.go
```

## マイグレーション

```
$ sql-migrate up -dryrun
$ sql-migrate up
```

## 問題点

- users テーブルの name 以外のカラムを API 経由で操作できてしまう
