package main

import (
	"fmt"
)

func add(a int, arg...int) int {
	var sum int =a
	for i :=0 ; i<len(arg);i++{
		sum += arg[i]
	}
	return sum
}

func main() {
	ret := add(10,3,3)
	fmt.Println(ret)
}