//统计耗时
package main

import (
	"fmt"
	"time"
)

func test(){
	time.Sleep(time.Millisecond*100) //100微妙
	for i :=0 ; i<10000000 ; i++{
		if(i ==10000000){
			break
		}
	}
}

func main()  {
	start := time.Now().UnixNano()
	test()
	end := time.Now().UnixNano()

	fmt.Println((end-start)/1000)

}