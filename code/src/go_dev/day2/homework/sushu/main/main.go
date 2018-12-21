package main

import (
	"fmt"
	"strconv"
)

//取101-201之内的素数,这个例子是最终结果是错的
func calc(a int ,b int) string{
	var ret string
	for i := a ; i<=b ;i++{
		for j:= 2 ; j<=i;j++{
			if (i % j ==0){
				break
			}else {
				ret = ret + " " +fmt.Sprintf("%d",i)
				break
			}

		}
	}
	return ret
}
//阶乘结果之和相加
func jiecheng(a int) int  {
	//a = 3   1*2*3
	var n int = 1
	var m int = 0
	for i := 1 ; i<=a; i++{
		n = n * i
		m = m + n
	}
	return m
}
//打印水仙花
func flower(a int ,b int) {
        //153 1**3 5**3 3**3

	for i := a ; i <=b ; i++{  //取值 153
		var ret int = 0
		str := fmt.Sprintf("%d",i)   //153 格式化字符串 用于下面len和字符串拼、接、拆
		for i_len := 0 ; i_len < len(str) ; i_len++ {    //几位数循环几次
			s_num := str[i_len:i_len+1]       //取出1或者5或者3
			s_T_n ,_:= strconv.Atoi(s_num)    //字符串转成数字用于下面计算
			n := s_T_n * s_T_n * s_T_n
			ret = ret + n
		}
		if (ret == i){
			fmt.Println("水仙花:",i)
		}

	}
}

func main()  {
	result := calc(100,201)
	fmt.Println(result)
	res := jiecheng(5)
	fmt.Println("jiecheng:",res)
	flower(100,600)

}

