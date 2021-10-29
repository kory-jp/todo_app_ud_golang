package controllers

import (
	"log"
	"net/http"
	"todo_app_ud_golang/app/models"
)

func signup(w http.ResponseWriter, r *http.Request) {
	// ブラウザから送られてくるHTTPリクエストにより処理を分岐させる
	if r.Method == "GET" {
		generateHTML(w, nil, "layout", "public_navbar", "signup")
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
			Password: r.PostFormValue("possword"),
		}
		// ユーザー作成のメソッドを上記の構造体に対して実行
		if err := user.CreateUser(); err != nil {
			log.Println(err)
		}

		// 完了後のページ遷移
		http.Redirect(w, r, "/", 302)
	}
}
