package main

import "fmt"
var UserID int
var UserPwd string
func main() {
	var key int
	var loop bool = true
	for loop{
		fmt.Println("------------------FuckQQ-----------------------")
		fmt.Println("\t\t1:注册账户")
		fmt.Println("\t\t2:登录FuckQQ")
		fmt.Println("\t\t3:退出")
		fmt.Println("\t\t:请选择（1-3）：")
		fmt.Scanf("%d\n",&key)

		switch key {
		case 1:
			fmt.Println("注册用户")
		case 2:
			fmt.Println("登录")
			loop = false
		case 3:
			fmt.Println("退出")
			loop = false
		default:
			fmt.Println("输入有误，重新输入。。。")
		}
	}

	if key == 2{
		fmt.Println("输入用户名：")
		fmt.Scanf("%d\n",&UserID)
		fmt.Println("输入密码：")
		fmt.Scanf("%s\n",&UserPwd)
		err := login(UserID,UserPwd)
		if err !=nil{
			fmt.Println("登录失败")
		}
	}
}
