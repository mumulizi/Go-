package main
//拷贝图片电影等，但是拷贝文本文档出错
import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CopyFile(src string,dst string) error {
	sFileName ,err := os.Open(src)
	if err != nil{
		fmt.Println("open sFile fail,err:",err)
		return err
	}
	defer sFileName.Close()
	reader := bufio.NewReader(sFileName)


	dFileName,err := os.OpenFile(dst,os.O_CREATE|os.O_WRONLY,0666)
	if err != nil{
		fmt.Println("dFile creat or open fail,err:",err)
		return err
	}
	defer dFileName.Close()
	writer := bufio.NewWriter(dFileName)
	_,err =io.Copy(writer,reader)
	if err != nil{
		return nil
	}
	return err

}
func main() {
	var (
		src string
		dst string
	)
	fmt.Println("输入源：")
	fmt.Scanln(&src)
	fmt.Println("输入目标：")
	fmt.Scanln(&dst)
	err := CopyFile(src,dst)
	if err ==nil{
		fmt.Println("拷贝完成")
		return
	}
	fmt.Println("err:",err)
}

