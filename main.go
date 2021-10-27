package main

import (
	"fmt"
	"todo_app_ud_golang/app/models"
)

func main() {
	// 空の構造体のUserを呼び出し
	// u := &models.User{}
	// u.Name = "test2"
	// u.Email = "test2@example.com"
	// u.Password = "testtest"
	// fmt.Println(u)

	// u.CreateUser()

	// u, _ := models.GetUser(1)
	// fmt.Println(u)

	// u.Name = "test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// u.DeleteUser()
	// user, _ := models.GetUser(2)
	// fmt.Println(user)
	// user.CreateTodo("Todo5")

	// todo一覧取得
	// todos, _ := models.GetTodos()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// ユーザーに紐づいたtodo取得
	user2, _ := models.GetUser(2)
	todos, _ := user2.GetTodosByUser()
	for _, v := range todos {
		fmt.Println(v)
	}
}
