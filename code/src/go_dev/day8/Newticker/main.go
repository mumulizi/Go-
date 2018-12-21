package main

import (
	"fmt"
	"time"
)

func queryDB(ch chan int)  {
	time.Sleep(time.Second)
	ch <- 10
}


func main() {
	ch := make(chan int)

	go queryDB(ch)

	t := time.NewTicker(time.Second)

	select {
	case v:= <- ch:
		fmt.Println("res:",v)
	case <-t.C:
		fmt.Println("timeout")
	}
}
