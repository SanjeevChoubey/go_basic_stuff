package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

type cntry struct {
	Country string
	Capital string
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ff": firstfive,
}

func firstfive(s string) string {
	s = strings.TrimSpace(s)
	s = s[:5]
	return s
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}
func main() {
	c1 := cntry{
		Country: "India",
		Capital: "New Delhi",
	}
	c2 := cntry{
		Country: "China",
		Capital: "Bejing",
	}
	c3 := cntry{
		Country: "USA",
		Capital: "Washington",
	}
	countries := []cntry{c1, c2, c3}
	// nf, err := os.Create("functemplate.html")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	//err = tpl.Execute(nf, countries)
	err := tpl.Execute(os.Stdout, countries)
	if err != nil {
		log.Fatalln(err)
	}
}
