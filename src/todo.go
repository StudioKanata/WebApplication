package main

import (
	"html"
	"html/template"
	"net/http"
	"strings"
)

var todoLists = make(map[string][]string)

// セッションIDをキーにしてリストを取得する
func getTodoList(sessionId string) []string {

	todoList, ok := todoLists[sessionId]
	if !ok {
		todoList = []string{}
		todoLists[sessionId] = todoList
	}

	return todoList
}

// todoListを返す
func handleTodo(w http.ResponseWriter, r *http.Request) {

	sessionId, err := ensureSession(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	todoList := getTodoList(sessionId)
	t, _ := template.ParseFiles("templates/todo.html")
	t.Execute(w, todoList)

}

// セッション上のtodoListに追加する
func handleAdd(w http.ResponseWriter, r *http.Request) {

	sessionId, err := ensureSession(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	todoList := getTodoList(sessionId)

	r.ParseForm()
	todo := strings.TrimSpace(html.EscapeString(r.Form.Get("todo")))
	if todo != "" {
		todoLists[sessionId] = append(todoList, todo)
	}
	// リダイレクト
	http.Redirect(w, r, "/todo", http.StatusSeeOther)
}
