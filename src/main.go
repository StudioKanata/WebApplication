package main

import (
	"html/template"
	"log"
	"net/http"
)

var todoList []string

func handleTodo(w http.ResponseWriter, r *http.Request) {
	// 本来なら、戻り値にerrを受け取りハンドリングする必要がある
	t, _ := template.ParseFiles("templates/todo.html")
	t.Execute(w, todoList)
}

func handleAdd(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	todo := r.Form.Get("todo")
	todoList = append(todoList, todo)
	handleTodo(w, r)
}

func main() {
	todoList = append(todoList, "顔を洗う", "朝食を食べる", "歯を磨く")

	// HTML内から読み込むCSSファイルを返すため
	// http.FileServerに渡すパスを調整staticを除いてパスを検索
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/todo", handleTodo)

	http.HandleFunc("/add", handleAdd)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("failed to start server : ", err)
	}
}
