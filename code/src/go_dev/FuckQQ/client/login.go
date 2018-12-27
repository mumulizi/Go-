package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_dev/FuckQQ/common/message"
	"net"
)

func login(UserId int, UserPwd string)(err error){
	//fmt.Printf("UserId:%d,UserPwd:%s",UserID,UserPwd)
	//return nil
	conn ,err := net.Dial("tcp","127.0.0.1:8889")
	defer conn.Close()
	if err !=nil{
		fmt.Println("Client net.Dial err:",err)
		return
	}
	var mes message.Message
	mes.Type = message.LoginMesType
	var loginMes message.LoginMes
	loginMes.UserId = UserId
	loginMes.UserPwd = UserPwd

	//将loginMes序列化
	data ,err := json.Marshal(loginMes)
	if err !=nil{
		fmt.Println("json Marshal err:",err)
		return
	}
	mes.Data = string(data)

	data,err = json.Marshal(mes)
	if err !=nil{
		fmt.Println("json2 Marshal err:",err)
		return
	}

	var pkglen uint32
	pkglen = uint32(len(data))
	var buf [4]byte
	//encoding/binary实现了简单的数字与字节序列的转换以及变长值的编解码
	binary.BigEndian.PutUint32(buf[:4],pkglen)

	n,err := conn.Write(buf[:4])
	if n!=4 ||err !=nil{
		fmt.Println("conn Write err:",err)
		return
	}
	//fmt.Printf("客户端发送消息内容%v，长度%d",string(data),pkglen)

	//发送消息本身
	_,err = conn.Write(data)
	if err !=nil{
		fmt.Println("conn.write(date)err:",err)
		return
	}
	//休眠20s
	//time.Sleep(20*time.Second)
	//fmt.Println("sleep 20s")

	mes,err = readPkg(conn)
	if err !=nil{
		fmt.Println("readpkg err:",err)
		return
	}
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data),&loginResMes)
	if loginResMes.Code ==200{
		fmt.Println("登录成功")
	}else {
		fmt.Println(loginResMes.Error)
	}


	return
}
