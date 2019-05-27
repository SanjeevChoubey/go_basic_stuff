package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type countries struct {
	Country string
	Capital string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {
	india := countries{
		Country: "India",
		Capital: "New Delhi",
	}
	Nepal := countries{
		Country: "Nepal",
		Capital: "Kathmandu",
	}
	Bangladesh := countries{
		Country: "Bangaldesh",
		Capital: "Dhaka",
	}
	cntry := []countries{india, Nepal, Bangladesh}
	err := tpl.Execute(os.Stdout, cntry)
	if err != nil {
		log.Fatalln(err)
	}

}
