package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)


func main() {
	c,err := redis.Dial("tcp","172.16.100.60:6379")
	if err != nil{
		fmt.Println("conn redis fail:",err)
		return
	}
	defer c.Close()
}
