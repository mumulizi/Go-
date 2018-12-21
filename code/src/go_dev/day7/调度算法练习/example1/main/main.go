package main

import (
	"fmt"
	"go_dev/day7/调度算法练习/example1/blance"
	"math/rand"
	"os"
	"time"
)

//接口练习，实现一个负载均衡调度算法，支持随机轮训

func main() {
	//insts := make([]*blance.Instance,16)
	var insts []*blance.Instance
	for i :=0 ;i <16; i++{
		host := fmt.Sprintf("192.168.%d.%d",rand.Intn(254),rand.Intn(254))
		one := blance.NewInstance(host,8080)
		insts = append(insts,one)
	}
	var balancer blance.Blancer
	var conf = "random"
	if len(os.Args) > 1{
		if os.Args[1] == "random" || os.Args[1] == "roundrobin"{
			conf = os.Args[1]
		}else {
			fmt.Printf("%s is not match , blance is default random.\n",os.Args[1])
		}
	}
	if conf == "random" {
		balancer = &blance.RandomBlance{}
	}else if conf == "roundrobin"{
		balancer = &blance.RoundRobinBalance{}
	}
	//balancer := &blance.RandomBlance{}
	//balancer := &blance.RoundRobinBalance{}
	for {
		inst,err := balancer.Doblance(insts)
		if err !=nil{
			fmt.Println("do blace err,",err)
			continue
		}
		fmt.Println(inst)
		time.Sleep(time.Second)
	}
}