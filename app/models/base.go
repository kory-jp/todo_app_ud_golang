package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"todo_app_ud_golang/config"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

const (
	tableNameUser = "users"
	tableNameTodo = "todos"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.Dbname)
	if err != nil {
		log.Fatalln(err)
	}

	// %sには第二引数のtableNameUserに代入されたusersが代入される
	cmdU := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			uuid STRING NOT NULL UNIQUE,
			name STRING,
			email STRING,
			password STRING,
			created_at DATETIME
		)`, tableNameUser)

	Db.Exec(cmdU)

	cmdT := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			content TEXT,
			user_id INTEGER,
			created_at DATETIME
		)`, tableNameTodo)

	Db.Exec(cmdT)

}

// ユーザーにかかるユニークなID作成
func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

// パスワードをバイト形式に変換
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
