package main

import "fmt"

func swap(a *int, b *int)  {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	first := 100
	second := 200
	swap(&first,&second)
	fmt.Println("first",first)
	fmt.Println("second",second)
}
