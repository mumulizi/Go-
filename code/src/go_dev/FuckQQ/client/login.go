package main

import "fmt"

func login(UserID int, UserPwd string)(err error){
	fmt.Printf("UserID:%d,UserPwd:%s",UserID,UserPwd)
	return nil
}
