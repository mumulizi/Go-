package main

import "fmt"

func main() {
	var ch chan  int
	ch = make(chan int,10)
	for i :=0;i<10;i++{
		ch <- i
	}
	close(ch)

	for  {
		var b int
		b ,ok := <- ch
		if ok ==false{
			fmt.Println("chan close")
			break
		}
		fmt.Println(b)
	}
	/*
	//如果上面不加close，那么for循环得写成下面这样才不会出现死锁。
	//chan_sync例子里exitchan不用关闭就是因为for循环里已经写明了退出条件，循环8次后退出
	//有for range的时候记得close掉chan
		for n :=0 ; n<10; n++ {
		v :=<-ch
		fmt.Println(v)
	}
	 */

}
