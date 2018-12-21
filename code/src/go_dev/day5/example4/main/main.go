package main

import "fmt"

//二叉树


type Student struct {
	Name string
	Age int
	Score float32
	left *Student
	right *Student
}

func trans(root *Student) {
	if root == nil{
		return
	}
	fmt.Println(root)
	trans(root.left)
	trans(root.right)
}

func main() {
	/*
	var root Student   //这是一种初始化方法,这个时候trans(&root)，传进去的&root是个地址，这种初始化返回的值类型
	root.Name = "stu1"
	root.Age = 18
	root.Score = 100
	//不定义left和right就相当于等于nil
	*/
	var root *Student = new(Student)  //这是root的另一种初始化，返回的是引用类型，trans的时候传进去的不用加&
	root.Name = "stu1"
	root.Age = 18
	root.Score = 100

	var left1 Student  //用这种初始化方法是为了看带& 和不带&，底下带&
	left1.Name = "stu2"
	left1.Age = 18
	left1.Score = 100
	root.left =  &left1

	var right1 *Student =new(Student) //这是另一种初始化方法，底下不用带&
	right1.Name = "stu3"
	right1.Age = 18
	right1.Score =100
	root.right = right1

	var left2 *Student = &Student{ //第三种初始化方法，底下也不加&
		Name:"stu4",
		Age:90,
		Score:99,
	}
	left1.left = left2

	trans(root)

}


