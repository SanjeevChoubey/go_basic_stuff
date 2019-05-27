package main

import (
	"io"
	"net/http"
	"text/template"
)

func main() {
	http.Handle("/pics/", http.FileServer(http.Dir("public")))
	//http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "foo ran")
}
func dog(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("dog.gohtml")
	if err != nil {
		http.Error(w, "Page Not found", 404)
	}
	tpl.ExecuteTemplate(w, "dog.gohtml", nil)
}
