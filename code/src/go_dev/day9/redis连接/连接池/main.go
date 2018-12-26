package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//定义全局pool变量
var pool *redis.Pool

//初始化连接池，init当程序启动的时候
func init() {
	pool = &redis.Pool{
		MaxIdle:8,   //最大空闲连接数
		MaxActive:0,  //表示数据库的最大连接数，0表示没有限制
		IdleTimeout:100,  //最大空闲时间
		Dial: func() (redis.Conn, error) {  //初始化连接的代码，连接哪个IP的redis
			return redis.Dial("tcp","172.16.100.17:6379")
		},
	}
}

func main() {
	//取出一个连接
	conn := pool.Get()
	defer conn.Close()

	_,err := conn.Do("Set","name","我是你baba")
	if err != nil{
		fmt.Println("conn err:",err)
		return
	}
	//r ,err := conn.Do("Get","name")  //这样打出来的不是字符串，而是asscii码
	r ,err := redis.String(conn.Do("Get","name"))
	if err !=nil{
		fmt.Println("get err:",err)
		return
	}
	fmt.Println(r)
	//如果我们要从连接池里取出连接一定要保证连接池没有关闭
}



