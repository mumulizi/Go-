package main

import (
	"fmt"
	"strconv"
)
//反转
func reverse(str string) string {
	var ret string
	lenStr := len(str)
	for i :=0 ;i <lenStr; i++ {
		//ret = ret + str[lenStr-i-1:lenStr-i]
		ret = ret + fmt.Sprintf("%c",str[lenStr-i-1])
	}
	return ret
}

func main()  {
	str1 := "hello"
	str2 := "world"
	//str3 := str1 + " " +str2
	//fmt.Println(str3)
	str3 := fmt.Sprintf("%s %s",str1,str2)
	fmt.Println(str3)

	//substr :=str3[9:11]
	substr := str3[10]
	fmt.Println("substr:",fmt.Sprintf("%c",substr))
	result :=reverse(str3)
	fmt.Println("res:",result)
	num := 123
	var num1 string
	num1 = "123"
	s := fmt.Sprintf("%d",num)  //数字转成字符串
	r := s[0:1]
	b ,_:= strconv.Atoi(num1)   //字符串转数字，不需要的参数用_ 接收
        fmt.Println("s:",s)
	fmt.Println("r:",r)
	t := fmt.Sprintf("%T",r)
	t1 := fmt.Sprintf("%T",s)
	t2 := fmt.Sprintf("%T",num)
	t3 := fmt.Sprintf("%T",b)
	fmt.Println(t,t1,t2,t3)
}
