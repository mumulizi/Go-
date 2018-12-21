package main

import (
	"fmt"
)

func bsrot(b []int){
    for i :=0 ; i<len(b)-1; i++{
	    //for j := 0; j<len(b)-2;j++{
		    if b[i] < b[i+1] {
			    b[i], b[i+1] = b[i+1], b[i]

		    }
	    fmt.Println("b:",b)
	    //}
    }
}
func main() {
	a := [...]int{1,3,2,5,4,9,6}
	for i :=0 ;i<len(a); i++{
		bsrot(a[:])
	}

	fmt.Println(a)
}
