package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type countries struct {
	Country string
	Capital string
}

func main() {

	cnt := countries{
		Country: "USA",
		Capital: "Washington",
	}

	err := tpl.Execute(os.Stdout, cnt)
	if err != nil {
		log.Fatalln(err)
	}

}
