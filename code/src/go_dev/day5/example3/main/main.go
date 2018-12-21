package main

import (
	"fmt"
	"math/rand"
)
//链表

type Student struct {
	Name string
	Age int
	Score float32
	next *Student
}

func trans(p *Student) {

	for p !=nil {
		fmt.Println(*p)
		p = p.next
	}
}
func main() {
	var head Student
	head.Name = "obama"
	head.Age = 30
	head.Score = 60

	/*
	var stu1 Student
	stu1.Name = "Trump"
	stu1.Age = 40
	stu1.Score = 66
	head.next = &stu1

	var stu2 Student
	stu2.Name = "Trump2"
	stu2.Age = 40
	stu2.Score = 66
	stu1.next = &stu2

	trans(&head)
	*/

	var tail =&head
	for i :=0 ;i<3;i++{
		var stu = Student{
			Name:fmt.Sprintf("stu%d",i),
			Age: rand.Intn(100),
			Score: rand.Float32() *100,
		}
		//fmt.Println("stu----:",stu)
		tail.next = &stu
		//fmt.Printf("tail:%s.next=%s\n",tail,stu)
		tail = &stu
	}

	trans(&head)




}