package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Any code you want here in this function")
}
func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
