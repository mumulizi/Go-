package main

import (
	"fmt"
	"sort"
)

func testSlice() {
	var slice []int
	//var arr [5]int =[...]int{1,2,3,4,5}
	var arr  =[...]int{1,2,3,4,5}
	slice = arr[:]
	slice = slice[:len(slice)-1]
	slice = append(slice,88)
	fmt.Println(slice)

}
func modifyl(a []int)  {
	a[1]=100
}

func testSlice3()  {
	var a[5]int = [5]int{1,2,3,4,5}
	s := a[1:]
	fmt.Printf("s%p,c[1]%p\n",s,&a[1])
	fmt.Println(a)
	s[2]=100
	fmt.Println(a)
	fmt.Printf("s%p,c[1]%p\n",s,&a[1])
	s=append(s,10010)
	fmt.Println(a)
	fmt.Println(s)
	fmt.Printf("s%p,c[1]%p\n",s,&a[1])
	s[2]=98
	fmt.Println(a)
	fmt.Println(s)
	fmt.Printf("s%p,c[1]%p\n",s,&a[1])
}
func testSlice2()  {
	var b []int = []int{1,2,3,4,5,6}
	modifyl(b)
	b = append(b,20)
	fmt.Println("b:",b)
	//var c = [5]int{1,2,3,4,5}
	//s := c[1:3]
	//fmt.Println(s,c[1])
	//fmt.Printf("s%p,c[1]%p",s,&c[1])

}
func strToArry()  {
	s :="hello world"
	s1 := []rune(s)
	s1[1]='B'
	str := string(s1)
	fmt.Println(str)
}

func sortTest()  {
	var a = [...]int{3,4,6,31,3,6,2}
	sort.Ints(a[:])
	fmt.Println(a)
	var s = [...]string{"b","a","c","ab","A","aA","aa","AA","Aa","B"}
	sort.Strings(s[:])
	fmt.Println(s)
}

func searchTest() {
	var a = [...]int{2,3,1,4,5}
	sort.Ints(a[:])
	index := sort.SearchInts(a[:],2)
	fmt.Println(index)
}

func testcopy()  {
	var a []int = []int{1,2,3,4,5}
	b := make([]int,10)
	copy(b,a)
	b = append(b,100)
	a = append(a,10086)
	fmt.Println(b)
	fmt.Println(a)

}
func main() {
	testSlice()
	testSlice2()
	testSlice3()
	strToArry()
	sortTest()
	searchTest()
	testcopy()
}