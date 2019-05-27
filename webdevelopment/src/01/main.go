package main

func main(){
 name := "Sanjeev choubey"
 str := fmt.Sprint(" <!DOCTYPE html>
 <html lang="en"> 
 <head>
  <meta charset="UTF-8">
  <title>Hello World</title>
</head>
<body>
  <h1> Sanjeev Choubey</h1>
 </body>
</html>"
)
 //fmt.Println(whtml)
 nf,err := os.Create("index.html")
 if err != nil{
	 log.Fatal("Error creating file ", err)
 }
 defer nf.Close()
 io.Copy(nf, strings.NewReader(str))
}