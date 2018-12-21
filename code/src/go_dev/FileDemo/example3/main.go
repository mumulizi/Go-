package main

import (
	"fmt"
	"io/ioutil"
)

//一次性读取整个文件到内存呢
func main() {
	file := "d:/test.txt"
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("read file err:%v", err)
	}
	//fmt.Printf("%v", data) //输出的是byte数组
	fmt.Printf("%v", string(data))
}