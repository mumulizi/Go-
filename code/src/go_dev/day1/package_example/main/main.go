package main

import (
	"go_dev/day1/package_example/calc"
	"fmt"
)

func main(){
	sum := calc.Add(100,200)
	sub := calc.Sub(200,100)
	fmt.Println("sum:",sum,"sub:",sub)
}