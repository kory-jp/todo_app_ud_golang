package controllers

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"text/template"
	"todo_app_ud_golang/app/models"
	"todo_app_ud_golang/config"
)

// 共通レイアウトとURLテンプレート、データを整形して返却
// dataはusersやtodosなどの複数の構造体データを取得する必要があるのでinterface型
// filenamesは共通レイアウトやTopページその他の未確定数のviewsファイルを代入する必要があるので可変数引数を受け取れる(...string)
func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		// Sprintf=Goの静的型付けにおいて、stringと他の型をstring型として一緒に扱うことができるようにする
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	// Mustは、関数の呼び出しをラップして（* Template、error）を返し、エラーがnilでない場合はパニックを起こすヘルパーです。これは、次のような変数の初期化で使用することを目的としています。
	templates := template.Must(template.ParseFiles(files...))
	// ExecuteTemplateは、指定された名前を持つlayoutに関連付けられたテンプレートを指定されたデータオブジェクトに適用し、出力をwrに書き込みます。
	templates.ExecuteTemplate(w, "layout", data)
}

func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("Invalid session")
		}
	}
	return sess, err
}

// 正規表現の型を定義
var validPath = regexp.MustCompile("^/todos/(edit|update|delete)/([0-9]+)$")

// URLに含まれているtodoIDを解析して取得、戻り値として返却
func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 正規表現に一致したURLであるか確認
		q := validPath.FindStringSubmatch(r.URL.Path)
		if q == nil {
			http.NotFound(w, r)
			return
		}
		// [0]edit/[1](edit|update)/[2]IDのID部分を抽出、string変換返却
		qi, err := strconv.Atoi(q[2])
		if err != nil {
			http.NotFound(w, r)
			return
		}

		fn(w, r, qi)
	}
}

func StartMainServer() error {
	// Config.goのStatic(静的ファイル)を定義
	files := http.FileServer(http.Dir(config.Config.Static))

	// web上で静的ファイルを読み込む際に/static/以下にCSS/JSファイルを読み込もうするが、
	// 現在のディレクトリ構成において静的ファイルはViews以下に配置されているので
	// web上でアクセスする際のURLリクエストから/static/の部分を取り除く(StripPrefix)。
	// 第二引数にて上記で定義した静的ファイルを渡しWEB上で表示
	http.Handle("/static/", http.StripPrefix("/static/", files))

	// URL登録とroute_main.goのhandlerとの結び付け/railsのroute.rb
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	// "/authenticate"=viewsのformのactionを参照
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/todos", index)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/todos/new", todoNew)
	http.HandleFunc("/todos/save", todoSave)
	// edit"/"のように最後に/をつけると/の部分のみ完全一致していることを求められる(/が末尾についてないとURL全体が完全一致を求められる)
	// parseURL(todoEdit)は初めにparseURLにてIDを取得して、そのIDをもとにtodoEditページを作成、遷移のような流れになる(ハンドラ関数のチェイン)
	http.HandleFunc("/todos/edit/", parseURL(todoEdit))
	http.HandleFunc("/todos/update/", parseURL(todoUpdate))
	http.HandleFunc("/todos/delete/", parseURL(todoDelete))
	// サーバーの起動
	// 第二引数にnilを渡すと存在しないページにアクセスすると404ページを表示する
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
