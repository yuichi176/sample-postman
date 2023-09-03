package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("サーバの起動に失敗しました:", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}
