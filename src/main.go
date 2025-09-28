package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// HTML内から読み込むCSSファイルを返すため
	// http.FileServerに渡すパスを調整staticを除いてパスを検索
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/todo", handleTodo)
	http.HandleFunc("/add", handleAdd)

	port := getPortNumber()
	fmt.Printf("listening port : %d\n", port)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("failed to start : ", err)
	}
}
