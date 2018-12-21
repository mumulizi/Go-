package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
	NickName string
	Age int
	Phone string
}

func testStruct()  {
	user1 := &User{
		Name:"你大爷",
		NickName:"你二大爷",
		Age:19,
		Phone:"13122222121",
	}

	data,err := json.Marshal(user1)
	if err != nil{
		fmt.Println("json fail,err:",err)
		return
	}

	fmt.Printf("%T",string(data))
}

func testMap()  {
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["Name"] = "user02"
	m["NiceName"] = "你大爷"
	m["age"] = 16

	data,err := json.Marshal(m)
	if err != nil{
		fmt.Println("map json err:",err)
		return
	}
	fmt.Printf("%s",data)
}

func testSlice()  {
	var m map[string]interface{}
	var s []map[string]interface{}

	m = make(map[string]interface{})
	m["Name"] = "user03"
	m["Age"] = 20
	m["id"] = 1

	s = append(s,m)


	m = make(map[string]interface{})
	m["Name"] = "user04"
	m["Passwd"] = "12345"

	s = append(s,m)
fmt.Println(s)
	data ,err := json.Marshal(s)
	if err !=nil{
		fmt.Printf("slice json fail ,err:",err)
		return
	}
	fmt.Printf("%s",data)
}


func main() {
	testStruct()
	//testMap()
	//testSlice()
}

