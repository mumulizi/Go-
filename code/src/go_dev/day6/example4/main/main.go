package main
//类型断言
import "fmt"

type TestInterface interface {
}


func Test(a TestInterface )  {
	b,ok :=a.(int)
	if ok ==false{
		fmt.Println("convert fail")
		return
	}
	b += 3
	fmt.Println(b)

}


//类型断言的另一种写法
func just(items ... interface{}) {
	for i,v := range items{
		switch v.(type) {
		case bool:
			fmt.Printf("%d is bool, value is %v\n",i,v)
		case int,int32,int64,int8:
			fmt.Printf("%d is int, value is %v\n",i,v)
		case string:
			fmt.Printf("%d is string, value is %v\n",i,v)
		}
	}
}


func main() {
	var b int
	Test(b)

	just("nihao",111,true)
}