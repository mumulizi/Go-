package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)
//读取文件显示在终端，带缓冲区的方式
func main()  {
	file,err := os.Open("D:/test.txt")
	if err !=nil{
		fmt.Println("open file err:",err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		str,err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Printf("read file end")
			return
		}
		fmt.Printf(str)
	}
	fmt.Println("file read end close")
}
