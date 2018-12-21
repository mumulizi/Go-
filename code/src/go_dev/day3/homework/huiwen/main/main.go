package main

import "fmt"
//判断是否为回文 abcba

func huiwen(s string) bool {
        s_len := len(s)
        var flag bool
	flag = false
	var r string
	for i :=0 ; i<s_len; i++{
		r = r + s[s_len-i-1:s_len-i]
	}
	if (r==s){
		flag = true
	}
	return flag
}

func main()  {
	var str string
	//for {
		fmt.Scanf("%s\n",&str)
                ret := huiwen(str)
		fmt.Println(ret)
	//}



}