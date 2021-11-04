# todo_app_ud_golang

## ローカル上での起動

```
git clone
```

ルートディレクトリへ移動

```
cd todo_app_ud_golang
```

modules 管理

```
go mod init
```

パッケージインストール

```
go get github.com/google/uuid v1.3.0
go get github.com/go-sql-driver/mysql
go get gopkg.in/go-ini/ini.v1
```

データベース設定

- パスワード設定無しの root@localhost ユーザー作成
- todo_app_ud_golang データベースを作成

コマンド実行

```
go run main.go
```

## lec 注意

### lec137

`app/views/template/public_navbar`において以下の{{}}に空白が含まれていないか注意

```
{{define "navbar"}}
```
