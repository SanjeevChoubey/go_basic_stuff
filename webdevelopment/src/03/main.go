package main

import (
    "log"
    "text/template"
    "os"
)
var tpl *template.Template

func init(){
    tpl = template.Must(template.ParseGlob("template/*"))
}
func main(){
    //tpl, err := template.ParseFiles("tpl.gohtml")
    // if err != nil{
    //     log.Fatalln(err)
    // }
    // err = tpl.Execute(os.Stdout, nil)
    // if err != nil{
    //     log.Fatalln(err)
    // }

// tpl, err = tpl.ParseFiles("tpl1.gohtml")
// if err != nil{
//         log.Fatalln(err)
//     }

    err := tpl.ExecuteTemplate(os.Stdout,"tpl.gohtml",nil)
    if err != nil{
        log.Fatalln(err)
    }

    err = tpl.ExecuteTemplate(os.Stdout,"tpl1.gohtml",nil)
    if err != nil{
        log.Fatalln(err)
    }
}