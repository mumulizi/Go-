package main

import "fmt"

//素数判断
func sushu(a int,b int)  {
	var ret string
	for n :=a ; n<=b; n++{
		var flag bool = true
		for i :=2 ; i < n; i++{
			if (n % i ==0){
				flag = false
				break
			}
		}
		if flag == true{
			ret = ret + " " + fmt.Sprintf("%d",n)
		}
	}
	fmt.Println(ret)
}


//阶乘
func jiecheng(a int)  {
	//5x4x3x2x1
		var sum int

		var n int =1
		for i:=1;i<=a;i++{
			n = n*i
			sum = sum + n
		}
		fmt.Println(n)


	fmt.Println(sum)
}
func main() {
	sushu(101,201)
	jiecheng(5)
}