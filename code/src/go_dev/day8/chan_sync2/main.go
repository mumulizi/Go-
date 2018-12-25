package main

import (
	"fmt"
)

func send(ch chan int,exitChanl chan bool)  {
	for i :=0 ;i<10 ;i++{
		ch <- i
	}
	exitChanl <- true
}
func recv(ch chan int,exitChanl chan bool)  {
	for i:=0;i<10;i++{
		v,ok := <-ch
		if !ok{
			break
		}
		fmt.Println(v)
	}
	exitChanl <- true

}
func main() {
	ch :=make(chan int,10)
	exitChan := make(chan bool,2)
	go send(ch,exitChan)
	go recv(ch,exitChan)

	for i:=0 ; i<2 ; i++{
		<- exitChan
	}
	close(ch) //好像在这关不关都行
}

