package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Web application!")
}

func main() {
	// リクエストに対する処理の登録
	http.HandleFunc("/", hello)
	// HTTPサーバーを起動する関数で、ポート8080で受け付ける
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("failed to start : ", err)
	}
}
