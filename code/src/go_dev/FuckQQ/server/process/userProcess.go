package process

import (
	"encoding/json"
	"fmt"
	"go_dev/FuckQQ/common/message"
	"go_dev/FuckQQ/server/model"
	"go_dev/FuckQQ/server/utils"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}


func (this *UserProcess)ServerProcesslogin(mes *message.Message)(err error) {
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data),&loginMes)
	if err != nil{
		fmt.Println("json unmarsha err:",err)
		return
	}
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	var loginResMes message.LoginResMes

	//使用redis内的数据验证用户登录是否正确
	user,err :=model.MyUserDao.Login(loginMes.UserId,loginMes.UserPwd)
	if err !=nil{
		loginResMes.Code = 500
		loginResMes.Error = "用户不存在，请注册"
	}else {
		loginResMes.Code =200
		fmt.Println(user.UserName,"登录成功")
	}
	//if loginMes.UserId == 100 && loginMes.UserPwd =="123456"{
	//	loginResMes.Code =200
	//}else {
	//	loginResMes.Code = 500
	//	loginResMes.Error = "user not found ,please register"
	//}
	data ,err := json.Marshal(loginResMes)
	if err !=nil{
		fmt.Println("json unmarsha err:",err)
		return
	}
	resMes.Data = string(data)
	data ,err = json.Marshal(resMes)
	if err != nil{
		fmt.Println("json.marshal fail :",err)
		return
	}
	tf := &utils.Transfer{
		Conn:this.Conn,
	}
	err = tf.WritePkg(data)
	return
}

