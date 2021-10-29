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

type Session struct {
	ID        int
	UUID      string
	Email     string
	UserID    int
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

func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	cmd := `
		select
			*
		from
			users
		where
			email = ?
	`
	err = Db.QueryRow(cmd, email).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	return user, err
}

func (u *User) CreateSession() (session Session, err error) {
	session = Session{}
	cmd1 := `
		insert into sessions (
			uuid,
			email,
			user_id,
			created_at
		) values (?, ?, ?, ?)
	`
	_, err = Db.Exec(cmd1, createUUID(), u.Email, u.ID, time.Now())
	if err != nil {
		log.Println(err)
	}
	cmd2 := `
		select
			id,
			uuid,
			email,
			user_id,
			created_at
		from 
			sessions
		where
			user_id = ?
			and email = ?
	`
	err = Db.QueryRow(cmd2, u.ID, u.Email).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserID,
		&session.CreatedAt,
	)
	return session, err
}

func (sess *Session) CheckSession() (valid bool, err error) {
	cmd := `
		select
			*
		from
			sessions
		where
			uuid = ?
	`
	err = Db.QueryRow(cmd, sess.UUID).Scan(
		&sess.ID,
		&sess.UUID,
		&sess.Email,
		&sess.UserID,
		&sess.CreatedAt,
	)

	if err != nil {
		valid = false
		return
	}
	if sess.ID != 0 {
		valid = true
	}
	return valid, err
}
