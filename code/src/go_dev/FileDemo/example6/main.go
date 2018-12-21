package main
//判断文件是否存在
import (
	"fmt"
	"os"
)

func PathExists(path string)(bool,error) {
	_,err :=os.Stat(path)
	//if err ==nil{
	//	return true,nil
	//}
	if !os.IsNotExist(err){
		return true,nil
	}
	return false,err
}
func main() {
	var file string
	fmt.Println("input your fileOrfilePath:")
	fmt.Scanln(&file)
	_,err := PathExists(file)
	if err !=nil{
		fmt.Println("file is not exist")
		return
	}
	fmt.Println("file is exist...")

}
