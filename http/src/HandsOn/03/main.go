package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {

	http.Handle("/", http.HandlerFunc(mydefault))
	http.Handle("/dog", http.HandlerFunc(mydog))
	http.Handle("/me", http.HandlerFunc(itsme))
	http.ListenAndServe(":8080", nil)
}

func mydefault(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Default screen")
}
func mydog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Jenny")
}
func itsme(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(w, "index.gohtml", "Sanjeev Choubey")
	if err != nil {
		log.Fatalln(err)
	}
}
