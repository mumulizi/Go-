package main

import "fmt"

//方法

type Student struct {
	Name string
	Age int
	Score int
	Sex string
}
//把方法定义到了外面，相当于py类里的 self.xxx = xxx,这个self随便叫，我是为了理解方便写的self
func (self *Student) init(name string,age int,score int)  {
	self.Name = name
	self.Age =  age
	self.Score = score
	fmt.Println(self)
}

func (self Student) get() Student{
	return self
}

func main() {
	var stu Student
	stu.init("alex",18,100)  //调用方法

	s :=stu.get()   //调用方法
	fmt.Println(s)
}