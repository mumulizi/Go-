package main
//文件内字母数字空格等的统计
import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CountNum struct {
	CharCount int
	NumCount int
	SpaceCount int
	DefaultCount int
}

func main()  {
	var count CountNum
	fileName := "d:/test.log"
	file ,err :=os.Open(fileName)
	if err != nil{
		fmt.Println("open file err:",err)
		return
	}
	reader := bufio.NewReader(file)
	defer file.Close()
	for{
		s1,err := reader.ReadString('\n')
		str :=[]rune(s1) //中文
		//fmt.Println(str)
		if err ==io.EOF{
			fmt.Printf("file read end\n")
			break
		}
		for _,v := range str{
			switch  {
			case v >= 'a' && v<='z':
				fallthrough
			case v>='A' && v<= 'Z':
				count.CharCount++
			case v>= '0' && v<='9':
				count.NumCount++
			case v == '\t' || v==' ':
				count.SpaceCount++
			default:
				count.DefaultCount++
			}
		}
	}
	fmt.Printf("字母个数：%v,数字个数：%v,空格个数:%v,其他：%v",count.CharCount,count.NumCount,count.SpaceCount,count.DefaultCount)
}
