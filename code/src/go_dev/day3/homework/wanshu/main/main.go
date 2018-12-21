package main
//100内完数
import (
	"fmt"
)

func wanshu(b int) string {
	//6 = 1 + 2 + 3
	var num string
	for a :=1 ; a<b ;a++ {
		var ret int
		for i := 1; i < a; i++ {
			if a % i == 0 {
				ret = ret + i
			}
		}
		if ret == a {
			num = num +" " +fmt.Sprintf("%d",a)
		}
	}
	return num
}

func main()  {
	ret := wanshu(100)
	fmt.Println(ret)
}