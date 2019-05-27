package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/dog.jpg", dogpic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html:charset=utf-8")
	io.WriteString(w, `<img src="/dog.jpg">`)
}

func dogpic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dog.jpg")
}
