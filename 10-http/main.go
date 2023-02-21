package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, HTTPサーバ")
}
func main() {
	http.HandleFunc("/", handler)
	// nilで省略した場合はhttp.HandleFuncなどで登録したハンドラが使用される
	http.ListenAndServe(":8080", nil)
}
