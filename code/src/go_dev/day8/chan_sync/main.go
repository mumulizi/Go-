package main

import (
	"fmt"
	"time"
)

func calcChan(taskChan,resChan chan int,exitChan chan bool)  {
	for v := range taskChan{
		flag := true
		for i := 2 ; i<v; i++{
			if v%i == 0{
				flag = false
				break
			}
		}
		if flag{
			resChan  <- v
		}
	}
	fmt.Println("exit")
	exitChan <- true
}

func main() {
	t1 :=  time.Now()
	intChan := make(chan int,100000)
	resultChan := make(chan int,100000)
	exitChan := make(chan bool,8)

	go func() {
		for i :=0 ;i<100000; i++{
			intChan <- i
		}
	close(intChan)
	}()
	for i :=0 ;i<8 ; i++{
		go calcChan(intChan,resultChan,exitChan)
	}

	//等待所有计算的goroute全部退出
	for i :=0 ; i<8 ; i++{
		<- exitChan
		fmt.Println("waiting...",i," exit")
	}
	close(resultChan)

	for v:= range resultChan{
		fmt.Println(v)
	}
t2 := time.Now()
fmt.Println(t2.Sub(t1))

}
