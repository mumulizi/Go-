package main

import (
	"fmt"
)

func main() {
	str := "hello world 中国"

	for i ,v := range str{
		//fmt.Printf("%T\n",v)
		if i>2{
			continue
		}
		if(i>3){
			break
		}
		fmt.Printf("i[%d], v[%c] len[%d]\n",i,v,len([]byte(string(v))))
	}
	n := (5%10) + '0'
	fmt.Println(n)

}
