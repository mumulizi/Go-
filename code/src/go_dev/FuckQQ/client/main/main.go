package main

import (
	"fmt"
	"go_dev/FuckQQ/client/process"
	"os"
)
var UserId int
var UserPwd string
func main() {
	var key int
	for {
		fmt.Println("------------------FuckQQ-----------------------")
		fmt.Println("\t\t1:注册账户")
		fmt.Println("\t\t2:登录FuckQQ")
		fmt.Println("\t\t3:退出")
		fmt.Println("\t\t请选择（1-3）：")
		fmt.Scanf("%d\n",&key)

		switch key {
		case 1:
			fmt.Println("注册用户")
		case 2:
			fmt.Println("登录")
			fmt.Println("输入用户名：")
			fmt.Scanf("%d\n",&UserId)
			fmt.Println("输入密码：")
			fmt.Scanf("%s\n",&UserPwd)
			up := &process.UserProcess{}
			up.Login(UserId,UserPwd)
			for {
				process.ShowMenu()
			}
		case 3:
			fmt.Println("退出")
		os.Exit(0)
		default:
			fmt.Println("输入有误，重新输入。。。")
		}
	}


}
