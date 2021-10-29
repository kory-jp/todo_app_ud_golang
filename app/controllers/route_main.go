package controllers

import (
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	// 引数のテンプレートファイルをParsFilesにて分析、テンプレートの構造体を作成
	//server.goのgenerateHTMLに以下のコードを含め、共通化
	// t, err := template.ParseFiles("app/views/templates/top.html")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// Executeにてテンプレートファイルを実行,第二引数にはデータの引き渡し
	// t.Execute(w, "Hello")
	_, err := session(w, r)
	if err != nil {
		// 引数には表示するHTMLパーツ、データを一覧揃える
		generateHTML(w, "Hello", "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/todos", 302)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	// sessionを取得できればindex,sessionが返ってこなければTOP
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "index")
	}
}
