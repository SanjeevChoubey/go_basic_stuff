package main
import(
	"fmt"
)
type stringmap map[string]string

func (c stringmap) print(){
	for country,capital := range c{
		fmt.Println("Capital of ", country ,"is ",capital)
	}
}

func main(){
	
	country := make(stringmap)
	country["india"] = "New delhi"
	country["pakistan"] = "islamabad"
	country["sri lanka"] = "colombo"
	//Print(country)
	country.print()
}


