//数组操作
package main

import (
	"fmt"
)

func test1() {
	var a [10]int
	fmt.Println(a)  //默认不赋值的时候是0,下标是从0开始

	a[1] = 200
	fmt.Println(a)

	for i:=0 ;i <len(a); i++{
		fmt.Println(a[i])
	}

	for i,v :=range a{
		fmt.Println(i,v)
	}
}

func test2() {
	var a[5]int
	fmt.Println(a)
	b:=a
	b[0]=100
	fmt.Println(b)
}

func test3(arr *[5]int)  {
	(*arr)[0] =100
	fmt.Println(*arr)
}

//二维数组
func test4()  {
	var a[2][5]int =[...][5]int{{1,2,3,4,5},{6,7,8,9,10}}
	fmt.Println(a)
	for row,v := range a{
		for col,v1 := range v{
			fmt.Printf("(%d,%d)=%d ",row,col,v1)
		}
		fmt.Println()
	}
}
func main()  {
	test1()
	test2()
	var a[5]int
	test3(&a)
	fmt.Println(a)
	test4()
}