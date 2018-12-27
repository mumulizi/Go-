package main

import (
	"fmt"
	"go_dev/FuckQQ/common/message"
	process2 "go_dev/FuckQQ/server/process"
	"go_dev/FuckQQ/server/utils"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}
func (this *Processor)serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		up := &process2.UserProcess{
			Conn:this.Conn,
		}
		err = up.ServerProcesslogin(mes)
	case message.RegisterMesType:
		//处理注册函数，没写
	default:
		fmt.Println("消息类型不存在，无法处理")
	}
	return
}

func (this *Processor)process()(err error){
	//defer conn.Close()
	//循环接收客户端的消息
	for{
		tf := &utils.Transfer{
			Conn:this.Conn,
		}
		mes,err := tf.ReadPkg()
		if err !=nil{
			if err  == io.EOF{
				fmt.Println("客户端退出，服务端也退出")
				return err
			}else {
				fmt.Println("readPkg err:",err)
				return err
			}
		}
		//fmt.Println("mes :",mes)

		err = this.serverProcessMes(&mes)
		if err != nil{
			return err
		}
	}

}