package model

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//我们在服务器启动后，就初始化一个userDao实例
//把他做成全局的变量，在需要和redis操作时，就直接使用即可
var (
	MyUserDao *UserDao
)

//定义一个userdao结构体
//完成对User结构体的各种操作
type UserDao struct {
	pool *redis.Pool
}
//使用工厂模式。创建一个userDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool:pool,
	}
	return
}

//根据用户id返回一个user实例+err
func (this *UserDao)getUserById(conn redis.Conn,id int)(user *User,err error) {
	//通过给定的id去redis查询这个用户
	res,err := redis.String(conn.Do("HGet","users",id))
	if err !=nil{
		if err == redis.ErrNil{   //在redis没有找到对应的
			err = ERROR_USER_NOEXISTS
		}
		return
	}
	//这里我们需要把res反序列化成user实例,上面返回的是指针
	user = &User{}
	err = json.Unmarshal([]byte(res),user)
	if err !=nil{
		fmt.Println("json ummarsh err=",err)
		return
	}
	return
}


//用户登录校验
func (this *UserDao)Login(userId int,userPwd string)(user *User,err error)  {
	//先从UserDao的连接池中取出一根连接
	conn := this.pool.Get()
	defer conn.Close()

	user,err =this.getUserById(conn,userId)
	if err !=nil{
		return
	}
	//这时证明用户存在获取到了，接下来判断密码是否正确
	if user.UserPwd != userPwd{
		err = ERROR_USER_PWD
		return
	}
	return
}