package main

import (
	"math/rand"
	"fmt"
)
func main() {
	for i :=0 ; i<10 ; i++{
	a := rand.Intn(10)
	fmt.Println(a)
	}
}