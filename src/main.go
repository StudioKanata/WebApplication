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

func main() {
	todoList = append(todoList, "顔を洗う", "朝食を食べる", "歯を磨く")

	// HTML内から読み込むCSSファイルを返すため
	// http.FileServerに渡すパスを調整staticを除いてパスを検索
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handleTodo)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("failed to start server : ", err)
	}

	/**
	  func hello(w http.ResponseWriter, req *http.Request) {
	  	fmt.Fprint(w, "Hello, Web application!")
	  }
	*/

	//func main() {
	//	// リクエストに対する処理の登録
	//	// 指定したdirectoryのファイルを返す
	//	http.Handle("/", http.FileServer(http.Dir("static/")))
	//	// HTTPサーバーを起動する関数で、ポート8080で受け付ける
	//	err := http.ListenAndServe(":8080", nil)
	//	if err != nil {
	//		log.Fatal("failed to start : ", err)
	//	}
}
