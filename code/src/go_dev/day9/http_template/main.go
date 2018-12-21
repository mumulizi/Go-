package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var myTemplate *template.Template
type Person struct {
	Name string
	Age int
	Title string
}


func userInfo(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("handle hello")
	//fmt.Fprint(w,"hello")

	p := Person{Name:"alex",Title:"MyTitle",Age:19}
	myTemplate.Execute(w,p)
}

func initTemplate(fileName string)(err error)  {
	myTemplate ,err = template.ParseFiles(fileName)
	if err != nil{
		fmt.Println("parse file err:",err)
		return
	}
	return
}


func main() {
	initTemplate("./index.html")
	http.HandleFunc("/user/info",userInfo)
	err := http.ListenAndServe("0.0.0.0:8888",nil)
	if err != nil{
		fmt.Println("http listen failed")
	}
}