package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", mydefault)
	http.HandleFunc("/dog", mydog)
	http.HandleFunc("/me", itsme)
	http.ListenAndServe(":8080", nil)
}

func mydefault(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "default")
}

func mydog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Jenny")
}
func itsme(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello sanjeev")
}
