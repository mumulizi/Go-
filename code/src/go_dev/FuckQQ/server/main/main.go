package main

import (
	"fmt"
	"net"
)
//func writePkg(conn net.Conn,data []byte) (err error){
//	var pkgLen uint32
//	pkgLen = uint32(len(data))
//	var buf [4]byte
//	binary.BigEndian.PutUint32(buf[:4],pkgLen)
//	n,err := conn.Write(buf[:4])
//	if n!=4 || err !=nil{
//		fmt.Println("conn.write(bytes)fail err:",err)
//		return
//	}
//	n,err = conn.Write(data)
//	if n!= int(pkgLen) || err !=nil{
//		fmt.Println("conn.write(bytes)fail err:",err)
//		return
//	}
//return
//}
//
//func readPkg(conn net.Conn) (mes message.Message,err error) {
//	buf := make([]byte,8096)
//	fmt.Println("读取客户端发送的数据")
//	_,err =conn.Read(buf[:4])
//	if err !=nil{
//		//fmt.Println("conn.read err:",err)
//		return
//	}
//	var pkgLen uint32
//	pkgLen = binary.BigEndian.Uint32(buf[:4])
//
//	n,err :=conn.Read(buf[:pkgLen])
//	if n!=int(pkgLen) ||err !=nil{
//		return
//	}
//	err = json.Unmarshal(buf[:pkgLen],&mes)
//	if err !=nil{
//		fmt.Println("json unmarshal fial err:",err)
//		return
//	}
//	return
//}

//func serverProcesslogin(conn net.Conn,mes *message.Message)(err error) {
//	var loginMes message.LoginMes
//	err = json.Unmarshal([]byte(mes.Data),&loginMes)
//	if err != nil{
//		fmt.Println("json unmarsha err:",err)
//		return
//	}
//	var resMes message.Message
//	resMes.Type = message.LoginResMesType
//	var loginResMes message.LoginResMes
//	if loginMes.UserId == 100 && loginMes.UserPwd =="123456"{
//		loginResMes.Code =200
//	}else {
//		loginResMes.Code = 500
//		loginResMes.Error = "user not found ,please register"
//	}
//	data ,err := json.Marshal(loginResMes)
//	if err !=nil{
//		fmt.Println("json unmarsha err:",err)
//		return
//	}
//	resMes.Data = string(data)
//	data ,err = json.Marshal(resMes)
//	if err != nil{
//		fmt.Println("json.marshal fail :",err)
//		return
//	}
//	err = writePkg(conn,data)
//	return
//}

//func serverProcessMes(conn net.Conn,mes *message.Message) (err error) {
//	switch mes.Type {
//	case message.LoginMesType:
//		err = serverProcesslogin(conn,mes)
//	case message.RegisterMesType:
//		//处理注册函数，没写
//	default:
//		fmt.Println("消息类型不存在，无法处理")
//	}
//	return
//}

//func process(conn net.Conn){
//	defer conn.Close()
//	//循环接收客户端的消息
//	for{
//		mes,err := readPkg(conn)
//		if err !=nil{
//			if err  == io.EOF{
//				fmt.Println("客户端退出，服务端也退出")
//				return
//			}else {
//				fmt.Println("readPkg err:",err)
//				return
//			}
//		}
//		//fmt.Println("mes :",mes)
//
//		err = serverProcessMes(conn,&mes)
//		if err != nil{
//			return
//		}
//	}
//}

func process(conn net.Conn) {
	defer conn.Close()
	processor := &Processor{
		Conn:conn,
	}
	err := processor.process()
	if err !=nil{
		fmt.Println("客户端和服务器通讯协程错误=err:",err)
		return
	}
}

func main() {
	fmt.Println("服务器开始监听8889端口")
	listen, err := net.Listen("tcp","0.0.0.0:8889")
	defer listen.Close()
	if err != nil{
		fmt.Println("net Dial err:",err)
		return
	}
	//一旦监听成功，就等待客户端连接
	for{
		conn,err := listen.Accept()
		if err !=nil{
			fmt.Println("Listen.Accept err:",err)
			continue
		}
		//连接成功启动一个协程和客户端通信
		go process(conn)
	}

}

