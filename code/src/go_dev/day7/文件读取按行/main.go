package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file ,err := os.Open("C:/test")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for  {
		data,prefix,err := reader.ReadLine()
		if err == io.EOF{
			break
		}

		fmt.Printf("data:%s,prefix:%v\n",string(data),prefix)
	}
}
