package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)


func main() {
	//连接redis
	c,err := redis.Dial("tcp","172.16.100.17:6379")
	if err != nil{
		fmt.Println("conn redis fail:",err)
		return
	}
	fmt.Println("connect success....",c)
	defer c.Close()
	//插入数据
	_,err = c.Do("set","name","大煞笔SB")
	if err != nil{
		fmt.Println("set err:",err)
		return
	}
	fmt.Println("set ok")
	//读取数据，返回的是字符串
	//r ,err := c.Do("Get","name")
	ret, err := redis.String(c.Do("Get","name"))
	if err !=nil{
		fmt.Println("get err:",err)
		return
	}
	//r := ret.(string) //转换出错，redis.string(c.Do("Get","name")) 才行
	//接口的转换是变量在前然后点类型，其他的转换是b = int64(a)
	//fmt.Println("ret:",r)
	fmt.Println("ret:",ret)


}
