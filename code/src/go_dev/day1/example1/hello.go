package main

import (
	"fmt"
	"time"
)

func nn(a int)  {
		fmt.Println(a)
}


func main() {
	//fmt.Println("hello world")
	for i := 0; i<100 ; i++ {
		go nn(i)
	}

	time.Sleep(1*time.Second)

}