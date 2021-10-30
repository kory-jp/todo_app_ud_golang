package controllers

import (
	"log"
	"net/http"
	"todo_app_ud_golang/app/models"
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
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		todos, _ := user.GetTodosByUser()
		user.Todos = todos
		generateHTML(w, user, "layout", "private_navbar", "index")
	}
}

func todoNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "todo_new")
	}
}

func todoSave(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		// フォームに入力されたデータを解析,パラメータを全て取得するためParseForm
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		// 上記で解析、取得したパラメータのうちPostFormValueのname項目のvalueを指定して取得
		context := r.PostFormValue("content")
		if err := user.CreateTodo(context); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/todos", 302)
	}
}

func todoEdit(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		t, err := models.GetTodo(id)
		if err != nil {
			log.Println(err)
		}
		generateHTML(w, t, "layout", "private_navbar", "todo_edit")
	}
}

func todoUpdate(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		content := r.PostFormValue("content")
		// $models.Todoにて空のTodoの構造体を作成
		t := &models.Todo{ID: id, Content: content, UserID: user.ID}
		// t.UpdateTodo()は上記の新規Todoを引数で受け取り、データベース上のデータを置き換える
		// 既存のデータを置き換えるのではなく、データを一から作成してそのデータを同じID上に上書きする
		if err := t.UpdateTodo(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/todos", 302)
	}
}
