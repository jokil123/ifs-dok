package main

import "net/http"

func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func main() {
	http.HandleFunc("/", test)
	http.ListenAndServe(":8080", nil)
}