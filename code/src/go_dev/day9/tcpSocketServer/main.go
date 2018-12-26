package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("start server....")
	listen,err := net.Listen("tcp","0.0.0.0:7000")
	if err != nil{
		fmt.Println("listen fail")
		return
	}
	for {
		fmt.Println("开始等待客户端连接")
		conn,err := listen.Accept()
		if err != nil{
			fmt.Println("accept fali,err-->:",err)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn)  {
	defer conn.Close()
	for {
		fmt.Println("开始接收客户端数据")
		buf := make([]byte,512)
		n ,err := conn.Read(buf)
		if err != nil{
			fmt.Println("read err:",err)
			return
		}
		fmt.Println("read:",string(buf[:n]))
	}
}


