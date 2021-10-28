package controllers

import (
	"net/http"
	"todo_app_ud_golang/config"
)

func StartMainServer() error {
	// Config.goのStatic(静的ファイル)を定義
	files := http.FileServer(http.Dir(config.Config.Static))

	// web上で静的ファイルを読み込む際に/static/以下にCSS/JSファイルを読み込もうするが、
	// 現在のディレクトリ構成において静的ファイルはViews以下に配置されているので
	// web上でアクセスする際のURLリクエストから/static/の部分を取り除く(StripPrefix)。
	// 第二引数にて上記で定義した静的ファイルを渡しWEB上で表示
	http.Handle("/static/", http.StripPrefix("/static/", files))

	// URL登録とroute_main.goのhandlerとの結び付け
	http.HandleFunc("/", top)

	// サーバーの起動
	// 第二引数にnilを渡すと存在しないページにアクセスすると404ページを表示する
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
