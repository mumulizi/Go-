package main

import (
	"encoding/json"
	"fmt"
)


type User struct {
	Name string `json:"name"`
	NickName string
	Age int
	Phone string
}


func testStruct() (ret string, err error) {
	user1 := &User{
		Name:"你大爷",
		NickName:"你二大爷",
		Age:19,
		Phone:"13122222121",
	}

	data,err := json.Marshal(user1)
	if err != nil{
		err = fmt.Errorf("json fail,err:",err)
		return
	}
	ret = string(data)
	fmt.Println(ret)
	return
}

func test() {
	data, err := testStruct()
	if err != nil{
		fmt.Println("errfail unjson")
		return
	}
	var user1 User
	err = json.Unmarshal([]byte(data),&user1)

	if err != nil{
		fmt.Println("2unjson fail",err)
		return
	}
	fmt.Println(user1)


}

func main() {
	test()
	testStruct()

}