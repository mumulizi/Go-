package main

import (
	"fmt"
	"os"
)

func main() {
	file,err := os.Open("D:/test.txt")
	if err !=nil{
		fmt.Println("err:",err)
		return
	}
	fmt.Printf("file:%v",file)

	err = file.Close()
	if err !=nil{
		fmt.Println("Close file err：",err)
	}
}
