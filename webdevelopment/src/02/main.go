package main
import(
    "log"
    "text/template"
    "os"
)

func main(){
	tpl,err := template.ParseFiles("tpl.gohtml")
    if err != nil{
        log.Fatalln(err)
    }
    
    nf, err := os.Create("index.html")
     if err != nil{
        log.Fatalln(err)
    }
    defer nf.Close()
    err1 := tpl.Execute(nf, nil)
     if err1 != nil{
        log.Fatalln(err1)
    }
}