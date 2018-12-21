package main

import (
	"fmt"
)

func main()  {
	var a int = 10
	fmt.Println(&a)    //打印a的内存地址

	var p *int         //p定义成指针
	p = &a             //p的指针内存地址，指到a的内存地址上
	*p =200            //把内存地址对应的值改成200
	fmt.Println(*p)    //取出指针p所指向的值
	fmt.Println(a)     //a的值也变了
}

