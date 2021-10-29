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

	generateHTML(w, "Hello", "layout", "top")
}
