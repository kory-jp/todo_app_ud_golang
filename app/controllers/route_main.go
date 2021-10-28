package controllers

import (
	"log"
	"net/http"
	"text/template"
)

func top(w http.ResponseWriter, r *http.Request) {
	// 引数のテンプレートファイルをParsFilesにて分析、テンプレートの構造体を作成
	t, err := template.ParseFiles("app/views/templates/top.html")
	if err != nil {
		log.Fatalln(err)
	}
	// Executeにてテンプレートファイルを実行,第二引数にはデータの引き渡し
	t.Execute(w, "Hello")
}
