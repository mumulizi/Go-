package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	var (
		CharCount int
		SpaceCount int
		NumCount int
		otherCount int
	)
	file,err := os.Open("C:/a")
	if err !=nil{
		fmt.Println("read file error",err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for  {
		str,err := reader.ReadString('\n')
		if err == io.EOF{
			break
		}
		if err != nil{
			fmt.Println("read str1 err,",err)
			break
		}
		runeArr := []rune(str)
		for _,v := range runeArr{
			switch  {
			case v>='a' && v <='z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				CharCount++
			case v ==' ' || v =='\t':
				SpaceCount++
			case v >='0' && v <='9':
				NumCount++
			default:
				otherCount++
			}

		}
	}
	fmt.Println("char count is %d",CharCount)
	fmt.Println("space count is %d",SpaceCount)
	fmt.Println("Num count is %d",NumCount)
	fmt.Println("other count is %d",otherCount)
}
