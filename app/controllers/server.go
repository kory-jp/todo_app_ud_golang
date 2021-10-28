package controllers

import (
	"net/http"
	"todo_app_ud_golang/config"
)

func StartMainServer() error {
	// URL登録とroute_main.goのhandlerとの結び付け
	http.HandleFunc("/", top)
	// サーバーの起動
	// 第二引数にnilを渡すと存在しないページにアクセスすると404ページを表示する
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
