package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", mydefault)
	http.HandleFunc("/dog", mydog)
	http.HandleFunc("/me", itsme)
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
