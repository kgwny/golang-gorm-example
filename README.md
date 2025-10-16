# golang-gorm-example

## go のインストール
```
brew install go
```

## go のバージョン/所在確認
```
go version
```
go version go1.24.4 darwin/arm64

```
which go
```
/opt/homebrew/bin/go

## go のバージョンアップ
必要に応じて実施すること
```
brew upgrade go
```

## go.mod の初期化
```
go mod init golang-gorm-example
```
go: creating new go.mod: module golang-gorm-example
go: to add module requirements and sums:
	go mod tidy

## echo のインストール
```
go get -u github.com/labstack/echo/v4
```

## gorm のインストール
```
go get -u gorm.io/gorm
```

## mysql ドライバーのインストール
```
go get -u gorm.io/driver/mysql 
```

## GORM の使い方

### CREATE
https://gorm.io/ja_JP/docs/create.html

### SELECT
https://gorm.io/ja_JP/docs/query.html

### 高度なクエリ
https://gorm.io/ja_JP/docs/advanced_query.html

### UPDATE
https://gorm.io/ja_JP/docs/update.html

### DELETE
https://gorm.io/ja_JP/docs/delete.html

### SQL Builder
https://gorm.io/ja_JP/docs/sql_builder.html
