package main

import (
	"fmt"
	"reflect"
)

//反射

type Student struct {
	Name string
	Age int
}

func Test(b interface{})  {
	t := reflect.TypeOf(b)
	fmt.Println(t)

	v := reflect.ValueOf(b)
	fmt.Println("v:",v)
	k := v.Kind()
	fmt.Println(k)
	iv := v.Interface()

	stu,ok := iv.(Student)
	if ok {
		fmt.Printf("%T\n\n",stu)
	}
}

func testInt(b interface{}) {
	val := reflect.ValueOf(b)
	val.Elem().SetInt(10)

	c := val.Elem().Int()
	fmt.Println(val,c)

}

func main() {
	//var a int = 200
	//Test(a)
	var a Student = Student{
		Name:"alex",
		Age:19,
	}
	Test(a)

	var b = 200
	testInt(&b)
	fmt.Println(b)
}
