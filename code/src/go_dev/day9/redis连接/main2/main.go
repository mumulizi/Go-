package main
//redis批量操作
import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	conn,err :=redis.Dial("tcp","172.16.100.17:6379")
	if err !=nil{
		fmt.Println("conn err:",err)
		return
	}
	defer conn.Close()
	_ , err = conn.Do("HMSet","user01","name","alex","age",20)
	if err !=nil{
		fmt.Println("HMSet err:",err)
		return
	}
	r,err := redis.Strings(conn.Do("HMGet","user01","name","age"))
	if err !=nil{
		fmt.Println("HMGet err:",err)
		return
	}
	fmt.Println(r)
	for i,v := range r{
		fmt.Printf("[%d]=%v",i,v)
	}
}