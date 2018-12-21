package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	Name string  `json:"student_name"`
	Age int
}

func main() {
	var stu student = student{
		Name:"alex",
		Age:10,
	}
	fmt.Println(stu)
	data,err := json.Marshal(stu)
	if err !=nil{
		fmt.Println("josn error:",err)
		return
	}else {
		fmt.Println(string(data))
	}
}

