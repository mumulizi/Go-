//输入一行 统计输入的字母、空格、数字、等出现的次数
package main

import (
	"fmt"
	"bufio"
	"os"
)

func count(str string)(WC,SC,NC,OC int) {
	t :=[]rune(str)   //rune 表示一个字符，中文在此处也算一个字符，不用他中文就是三个字符了
	for _,v := range t{
		switch  {
		case v >= 'a' && v <= 'z':
			fallthrough
		case v >= 'A' && v<='Z':
			WC++
		case v == ' ':
			SC++
		case v >='0' && v<='9':
			NC++
		default:
			OC++

		}
	}
	return
}

func main(){
	reader := bufio.NewReader(os.Stdin)
	result,_,err := reader.ReadLine()
	if err != nil {
		fmt.Println("read from console err",err)
		return
	}
	WC,SC,NC,OC :=count(string(result))
	fmt.Println(WC,SC,NC,OC)
}

