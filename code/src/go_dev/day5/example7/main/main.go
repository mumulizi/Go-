package main

import "fmt"

//接口

type Test interface {
	Print()
}

type Student struct {
	name string
	age int
	score int
}

func (p Student) Print()  {
	fmt.Println("name:",p.name)
	fmt.Println("age:",p.age)
	fmt.Println("score:",p.score)
}

func main()  {
	var t Test
	var stu Student = Student{
		name:"alex",
		age:10,
		score:90,
	}
	t = stu
	t.Print()
}
