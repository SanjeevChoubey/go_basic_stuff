package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/dog.jpg", dogpic)
	http.ListenAndServe(":8080", nil)

}

func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html:charset=utf-8")
	io.WriteString(w, `<img src="/dog.jpg">`)
}

func dogpic(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("dog.jpg")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		log.Fatalln(err)
	}
	http.ServeContent(w, r, fi.Name(), fi.ModTime(), f)
}
