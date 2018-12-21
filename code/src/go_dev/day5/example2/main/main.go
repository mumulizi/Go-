package main
//
import (
	"fmt"
)

type Student struct {
	Name string
	Age int
	score float32 //小写包外不能访问
}

func main() {
	//第一种初始化方式
	var stu Student
	stu.Age =18
	stu.Name = "alex"
	stu.score = 80

	//第二种初始化方法
	var stu1 *Student = &Student{
		Age:19,
		Name:"bob",
	}
	fmt.Println(*stu1)
	fmt.Println(stu1.Name)

	//第三种
	var stu3 *Student = new(Student)
	stu3.Age = 10086
	stu3.Name = "jeny"
	stu3.score = 1
	fmt.Println(*stu3)
	fmt.Println(stu3.Name)

	//第四种
	var stu4 = Student{
		Age:10010,
		Name:"linda",
		score:90,
	}
        fmt.Println(stu4)
        fmt.Println(stu4.Age)



	fmt.Println(stu)
	fmt.Printf("Name:%p\n",&stu.Name)
	fmt.Printf("Age:%p\n",&stu.Age)
	fmt.Printf("score:%p\n",&stu.score)
}

