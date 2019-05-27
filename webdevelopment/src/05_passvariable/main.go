package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template
var mydata struct {
	age int
}

func init() {
	tpl = template.Must(template.ParseGlob("template/*"))
}

func main() {
	mydata.age = 76
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", mydata.age)
	if err != nil {
		log.Fatal(err)
	}
}
