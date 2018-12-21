package main
//写文件
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "D:/test1.txt"
	file,err := os.OpenFile(filePath,os.O_CREATE|os.O_WRONLY,0666)
	if err !=nil{
		fmt.Printf("open file err %v",err)
		return
	}

	defer file.Close()

	str := "hello world\r\n"
	writer := bufio.NewWriter(file)
	for i:=0;i<5;i++{
		writer.WriteString(str)
	}
	writer.Flush()

}
