package main
//接口嵌套

import "fmt"

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type ReadWriter interface {
	Reader
	Writer
}

type File struct {
}

func (f *File) Read() {
	fmt.Println("Read")
}

func (f *File) Write() {
	fmt.Println("write")
}
func Test(rw ReadWriter)  {   //这里为什么是Readwrite类型，底下传进来的是Flie类型，ex4例子断言，接口可以转换，如果把main函数里的File改成int也行，只不过int没有实现ReadWriter接口的方法
	rw.Read()
	rw.Write()
}

func main() {
	//var f File
	//Test(&f)

	var f *File
	var b interface{}
	b=f
	v,ok := b.(ReadWriter)
	fmt.Println(v,ok)
}