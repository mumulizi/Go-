package main

import (
	"fmt"
)

func main()  {
	var a = 100
	fmt.Println("a=:",a)
	var b chan int = make(chan int,1)
	fmt.Println("b:",b)
}
