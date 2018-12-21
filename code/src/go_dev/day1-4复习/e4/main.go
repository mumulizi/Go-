package main

import "fmt"

func test(a int,b int) (c int) {
	if true{
		c :=a + b
		fmt.Println(c)
	}
	return
}
func main() {
	fmt.Println(test(10,20))
}
