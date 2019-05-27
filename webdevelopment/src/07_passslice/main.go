package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("template/*"))
}

func main() {
	country := []string{"india", "pakistan", "Sri Lanka"}
	//err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", country)
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.1.gohtml", country)
	if err != nil {
		log.Fatal(err)
	}

}
