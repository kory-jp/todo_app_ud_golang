package models

import (
	"log"
	"time"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

// main.goにて呼び出された構造体のUserのポインタを変数名uで引数で渡す
// メソッド内で構造体の値を書き換える必要があるためポインタメソッドを使用
func (u *User) CreateUser() (err error) {
	cmd := `
		insert into users (
			uuid,
			name,
			email,
			password,
			created_at
		) values (?, ?, ?, ?, ?)
	`
	// Exec = 単純にクエリを実行し、結果行を戻さないメソッドです。たいていは delete, insert で使うと思います。
	_, err = Db.Exec(
		cmd,
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.Password),
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `
		select
			id,
			uuid,
			name,
			email,
			password,
			created_at
		from
			users
		where
			id = ?
	`

	// QueryRow = QueryRow()はRow型を返します。一行のSQL結果が返る
	// 各行において、rows.Scan()でカラムをポインタ変数へ読み込んでいます。
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	return user, err
}

func (u *User) UpdateUser() (err error) {
	cmd := `
		update
			users
		set
			name = ?,
			email = ?
		where
			id = ?
	`
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (u *User) DeleteUser() (err error) {
	cmd := `
		delete
		from
			users
		where
			id = ?
	`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
