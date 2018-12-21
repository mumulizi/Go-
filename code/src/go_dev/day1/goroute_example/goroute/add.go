package goroute
//c chan int 定义c管道 类型数字
func Add(a int,b int,c chan int)  {
	sum := a + b
	c <- sum
}
