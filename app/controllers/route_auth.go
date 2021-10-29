package controllers

import (
	"log"
	"net/http"
	"todo_app_ud_golang/app/models"
)

func signup(w http.ResponseWriter, r *http.Request) {
	// ブラウザから送られてくるHTTPリクエストにより処理を分岐させる
	if r.Method == "GET" {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, nil, "layout", "public_navbar", "signup")
		} else {
			http.Redirect(w, r, "/todos", 302)
		}
	} else if r.Method == "POST" {
		// r.ParseForm=ブラウザから送られてきたデータの解析
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		// Userの構造体をインスタンスして、ブラウザからのデータをセット
		user := models.User{
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			Password: r.PostFormValue("password"),
		}
		// ユーザー作成のメソッドを上記の構造体に対して実行
		if err := user.CreateUser(); err != nil {
			log.Println(err)
		}

		// 完了後のページ遷移
		http.Redirect(w, r, "/", 302)
	}
}

// ログインページの表示
func login(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, nil, "layout", "public_navbar", "login")
	} else {
		http.Redirect(w, r, "/todos", 302)
	}
}

// ログイン機能実装
func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	// メールアドレスからユーザーを検索
	user, err := models.GetUserByEmail(r.PostFormValue("email"))
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/login", 302)
	}
	// パスワードの検証,データベースの暗号化されたパスワードとブラウザにて入力されたパスワードを暗号化して比較
	if user.Password == models.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			log.Println(err)
		}
		// cookieの設定
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		// cookieのセット
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}
