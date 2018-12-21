package main

import "fmt"

type AInterface interface {
	InterFaceSay(name string)
}
type Student struct {
	Name string
}
type Teacher struct {
	Name string
}
func (s Student) InterFaceSay(name string) {
	fmt.Printf("%s say\n",name)
}
func (t Teacher) InterFaceSay(name string) {
	fmt.Printf("%s say\n",name)
}
func Say(s string)  {
	fmt.Printf("%s say\n",s)
}
func main() {
	s1 := Student{
		Name:"stu1",
	}
	var a AInterface = s1
	a.InterFaceSay(s1.Name)

	t1 := Teacher{
		Name:"tea1",
	}
	var b AInterface = t1
	b.InterFaceSay(t1.Name)

	Say(s1.Name)
	Say(t1.Name)
}
