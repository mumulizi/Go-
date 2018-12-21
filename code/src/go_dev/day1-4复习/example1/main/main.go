package main

import (
	"fmt"
	"strings"
)
//字符串
func test1()  {
	var s string
	sLen := len(s)
	fmt.Printf("s:%s len is%d memAdd%p:\n",s,sLen,&s)
	s = s + "aca Bcd"
	sLen = len(s)
	fmt.Printf("s:%s len is%d memAdd%p:\n",s,len(s),&s)
        r := strings.HasPrefix(s,"http://")
	r2 := strings.HasSuffix(s,".com")
	i := strings.Index(s,"d")
	fmt.Println(i)
	r3 := strings.Contains(s,"e")
	r4 := strings.Contains(s,"d")
	fmt.Println(r3,r4)
	var s1 string
	r5 := strings.Repeat(s,1)
	r6 := strings.Count(r5,"a")

	r7 := strings.Replace(s,"a","www.",1)
	r8 := strings.Replace(r7,"c","o",-1)
	r9 := strings.ToLower(s)
	r10 := strings.ToUpper(s)
	r11 := strings.Fields(s)
	r12 := strings.Split(s,"c")
	r13 := strings.Join(r12,"")
	fmt.Println(r5,r6,r7,r8,r9,r10,r11,r12,r13)
	if !(r && r2) {
		s1 = fmt.Sprintf("http://%s.com",s)
	}

	fmt.Println(s1)
	var n int= 1
	var ip *int
	fmt.Printf("%p\n",ip)
	ip = &n
	fmt.Printf("%p\n",ip)
	fmt.Printf("%p\n",&ip)
	fmt.Printf("%p",&n)

}

func main() {
	test1()
}
