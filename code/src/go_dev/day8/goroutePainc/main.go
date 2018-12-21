package main
//recover 捕获goroutine错误
import (
	"fmt"
	"time"
)

func calc() {
	defer func() {
		err := recover()
		if err != nil{
			fmt.Println(err)
		}
	}()


	var p *int
	*p = 100
	fmt.Println(p)
}

func main() {
	go calc()
	time.Sleep(time.Second * 3)
	fmt.Println("eixt...")
}
