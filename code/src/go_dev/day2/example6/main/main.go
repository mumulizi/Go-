package main

import (
	"fmt"
	"time"
)
const (
	man =1
	femal =2
)
func main(){
	for  {
		second := time.Now().Unix()
		if (second % 2 ==0){
			fmt.Println("femal")
		}else {
			fmt.Println("man")
		}
		time.Sleep(1*time.Second)
	}


}