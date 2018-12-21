package main

import (
	"fmt"
)

func nine(n int) {
	var (
		i int
		j int
	)
	//1*1=1
	//1*1=1 2*2=4
	for i =1 ; i<=n; i++{
		for j =1 ;j <=i ;j++{
			fmt.Printf("%d*%d=%d,",j,i,i*j)
		}
		fmt.Println()
	}
}

func main() {
	nine(9)
}
