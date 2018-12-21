package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/test",Hello)  //相当于django的url
	err := http.ListenAndServe("0.0.0.0:8888",nil)
	if err != nil{
		fmt.Println("http listen fail")
	}
}

func Hello(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("handle hello")
	fmt.Fprint(w,"helloworld")
}

