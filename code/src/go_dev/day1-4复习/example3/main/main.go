package main

import "fmt"

//数组
func array() {
	var a [5]int
	fmt.Println("init a:", a)
	a[0] = 1
	a[4] = 5
	fmt.Println(a[:4])
	var b [3]int = [...]int{1, 2, 3}
	b[0] = 4
	fmt.Println(b)
	c := [...]int{9, 8, 7, 6}
	//for i:=0 ;i<len(c); i++{
	//	fmt.Println(c[i])
	//}
	for _, v := range c {
		fmt.Println(v)
	}
}

//切片
func slience(){
	a :=[10]int{1,2,3,4,5,6,7,8,9,0}
	b :=a[:]
	fmt.Println(b)
	b[0]=89
	fmt.Println(a)
	b=append(b,10086)
	fmt.Println(b)
	fmt.Println(a)

	d :=[]int{}
	fmt.Println(d)
	d=append(d,5)
	fmt.Println(d)
	d[0]=10086
	fmt.Println(d)

	c :=a[:]
	c1 :=a[:]
	fmt.Println("after c:",a)
	c[0] = 88
	fmt.Println("c0 c:",a)
	c1[1] = 77
	fmt.Println("c1 c:",a)




}
func main() {
	//array()
	slience()
}

