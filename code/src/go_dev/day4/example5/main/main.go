//map相关 字典相关
package main

import (
	"fmt"
)

func mapTest1()  {
	/*var a map[string]string = map[string]string{
		"123":"456",
	} //这是一种定义map的方式*/
	a := make(map[string]string,10) //这是第二种定义map方式。10是长度，可加可不加，会自动扩

	a["a"]="A"
	a["b"]="B"
	fmt.Println("a--->:",a)
	fmt.Println("a->v:",a["a"])
	delete(a,"a")   //删除某个key
	for k,v := range a{
		fmt.Println(k,v,"-----")
	}
}

//map里面套map。相当于字典里面套字典
func mapTest2() {
	a := make(map[string]map[string]string)
	a["k1"] = make(map[string]string)  //make初始化第二层的map，map定义后要make不然不能用
	a["k1"]["k1-1"] = "v1"
	a["k1"]["k1-2"] = "v2"
	a["k1"]["k1-3"] = "v3"
	fmt.Println(a)
}

func modify(a map[string]map[string]string) {
	//val,ok := a["zhangsan"]
	_,ok := a["zhangsan"]
	//fmt.Println("val:",val)
	fmt.Println("ok:",ok)
	/*if ok {
		val["pwd"]="123"
		val["id"]="GG"
	}else {
		a["zhangsan"] = make(map[string]string)
		a["zhangsan"]["pwd"]="123"
		a["zhangsan"]["id"]="GG"

	}*/
	if !ok{
		a["zhangsan"] = make(map[string]string)
	}
	a["zhangsan"]["pwd"]="123"
	a["zhangsan"]["id"]="GG"
	return
}
func mapTest3() {
	a := make(map[string]map[string]string)
	//a["zhangsan"] = make(map[string]string)
	modify(a)
	fmt.Println(a)
}


func main() {
	mapTest1()
	mapTest2()
	mapTest3()
}


