package main

import (
	"log"
	"os"
	"text/template"
)

type countries struct {
	Country string
	Capital string
}

type games struct {
	Player   string
	Strength string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	c1 := countries{
		Country: "India",
		Capital: "New Delhi",
	}
	c2 := countries{
		Country: "France",
		Capital: "Paris",
	}
	c3 := countries{
		Country: "China",
		Capital: "Bejing",
	}
	g1 := games{
		Player:   "Bhuwaneshwar kumar",
		Strength: "Swing",
	}

	g2 := games{
		Player:   "Federer",
		Strength: "Back hand",
	}
	Manycountries := []countries{c1, c2, c3}
	Manygames := []games{g1, g2}
	data := struct {
		Cntry []countries
		Game  []games
	}{
		Manycountries,
		Manygames,
	}
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
