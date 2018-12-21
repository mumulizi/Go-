package main

import (
	"fmt"
	"runtime"
)

func numCheng(x ,n int64) int64 {
   var r int64 = x
	for i := 1 ; int64(i)<n ; i++{
		r = x * r
	}
	return r
}


func test()  {
	for z := 0; int64(z) <=10000 ;z++{
		for i := 0;int64(i) <=900; i++{
			for in := 2 ; in <=4;in++{
				if  numCheng(int64(i),int64(in)) + numCheng(int64(900-i),int64(in)) == numCheng(int64(z),int64(in)){
					fmt.Printf("%d ^ %d + %d ^ %d = %d ^ %d \n",i,in,(900-i),in,z,in)
				}

			}
		}
	}

}

func main() {
	num := runtime.NumCPU()
	fmt.Println(num)
}