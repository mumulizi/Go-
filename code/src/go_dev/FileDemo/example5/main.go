package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//将一个文件的内容导入到另一个文件内

func WrFile(s string)  {
	filePath := "D:/test.log"
	file,err :=os.OpenFile(filePath,os.O_CREATE|os.O_APPEND,0666)
	if err !=nil{
		fmt.Printf("open writfile fail:%v",err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString(s)
	writer.Flush()


}
func main() {
	filePath := "D:/test.txt"
	file,err := os.Open(filePath)
	if err != nil{
		fmt.Println("open file err:",err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str,err := reader.ReadString('\n')
		if err ==io.EOF{
			fmt.Printf("err read fail %v",err)
			return
		}
		fmt.Println(str)
		WrFile(str)
	}

}
