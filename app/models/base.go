package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"todo_app_ud_golang/config"

	"github.com/google/uuid"
	// _ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

var err error

const (
	tableNameUser    = "users"
	tableNameTodo    = "todos"
	tableNameSession = "sessions"
)

func init() {
	// MySQLのDSM＝ username:password@address|ip|tcp|port/dbname
	// ?parseTime=trueはMySQL側のcreated_atとGolang側のCreatedAt(time.Time)の差異から生じるエラーを解消する
	DSN := fmt.Sprintf("%s:%s@%s/%s?parseTime=true", config.Config.UserName, config.Config.Password, config.Config.DBPort, config.Config.DBname)
	Db, err = sql.Open(config.Config.SQLDriver, DSN)
	if err != nil {
		log.Fatalln(err)
	}

	// 接続成功したかを確認
	err = Db.Ping()
	if err != nil {
		fmt.Println("データベース接続失敗")
		fmt.Println(err)
		return
	} else {
		fmt.Println("データベース接続成功")
	}

	// sqlite3とMySQLでSQL文に差異があるのでその部分を修正
	cmdU := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY auto_increment,
		uuid varchar(50) NOT NULL UNIQUE,
		name varchar(50),
		email varchar(50),
		password varchar(50),
		created_at datetime default current_timestamp
	)`, tableNameUser)

	_, err := Db.Exec(cmdU)
	if err != nil {
		log.Fatalln(err)
	}

	cmdT := fmt.Sprintf(`
	create table if not exists %s (
		id integer primary key auto_increment,
			content text,
			user_id integer,
			created_at datetime default current_timestamp
	)`, tableNameTodo)

	Db.Exec(cmdT)

	cmdS := fmt.Sprintf(`
	create table if not exists	%s (
		id integer primary key auto_increment,
			uuid varchar(50) not null unique,
			email varchar(50),
			user_id integer,
			created_at datetime default current_timestamp
	)`, tableNameSession)

	Db.Exec(cmdS)

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
