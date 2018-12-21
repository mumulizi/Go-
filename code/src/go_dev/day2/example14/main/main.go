package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	for i :=0 ; i<10 ;i++ {
		a := rand.Int()
		b := rand.Intn(100)
		c := rand.Float32()

		fmt.Println("a:",a)
		fmt.Println("b:",b)
		fmt.Println("c:",c)
	}
}