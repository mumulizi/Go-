package main

import "fmt"

type Car interface {
	GetName() string
	Run()
}

type Byd struct {
	Name string
}

//type Bmw struct {
//	Name string
//}

func (p *Byd) GetName() string {
	return p.Name
}

func (p *Byd) Run()  {
	fmt.Printf("%s running",p.Name)
}

func main() {
	var c Car
	b := &Byd{
		Name:"Tang",
	}

	c = b
	c.Run()

}



