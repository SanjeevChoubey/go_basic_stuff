package main

import (
	"fmt"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("template/*"))
}
func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 34)
	if err != nil {
		fmt.Println(err)
	}

}
