package main

import "fmt"

//继承

type Car struct {
	Name string
}

type Bike struct {
	Car
	speet int
}

type Train struct {
	t Car
}

func (self *Car) Run() {
	fmt.Printf("%s sousousouRuning...\n",*self)
}

//只要下面print实现了string那就会默认按照此格式打印
func (slef Train) String() string  {
	str := fmt.Sprintf("name:%s",slef.t.Name)
	return str
}

func main() {
	var c Bike
	c.speet = 20
	//继承字段
	c.Name = "ofo"
	fmt.Println(c)
    //继承方法
	c.Run()
	//组合
	var d Train
	d.t.Name = "huoche"
	d.t.Run()
	//发现%s，去找d有没有实现string接口，实现那就按照定义的格式打印
	fmt.Printf("%s",d)
}
