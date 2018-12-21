package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age int
	Score float32
	sex string
}

func (s Student) Print()  {
	fmt.Println("---->",s)
}
func (s Student) Print1()  {
	fmt.Println("----1>",s)
}
func (s Student) Set(name string)  {
	s.Name = name
}

func TestStruct(b interface{})  {
	//tye := reflect.TypeOf(b)
	val := reflect.ValueOf(b)
	kd := val.Kind()
	if kd != reflect.Ptr && val.Elem().Kind() ==reflect.Struct{
		fmt.Printf("not struct type")
		return
	}
	num := val.Elem().NumField()  //去获取多少个字段。Name Age Score sex
	val.Elem().Field(0).SetString("stu10086")  //修改字段的值
	for i :=0 ;i <num; i++{
		fmt.Printf("numfiled %d %v\n",i,val.Elem().Field(i).Kind()) //获取字段类型，去掉kind获取字段具体的值
		//fmt.Printf("numfiled %d %v\n",i,val.Elem().Field(i)) //获取字段类型，去掉kind获取字段具体的值
	}
	fmt.Printf("%d has filed",num)
	numMethod := val.Elem().NumMethod()  //去获取多少个方法。Set
	fmt.Printf("%d has method",numMethod)

	var params []reflect.Value
	val.Elem().Method(1).Call(params)    //调用结构体中的方法，1是第二个方法的下标。0是第一个

}


func main() {
	var stu Student = Student{
		Name:"alex",
		Age:19,
		Score:99.9,
	}

	TestStruct(&stu)

}
