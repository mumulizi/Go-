package main

import (
	"fmt"
	"strconv"
)

func main()  {
	var str string
	fmt.Scanf("%s",&str)

	ss, err := strconv.Atoi(str)
	if  err != nil{
		fmt.Println("can't convert to int",err)
		return
	}
	fmt.Println(ss)
}

