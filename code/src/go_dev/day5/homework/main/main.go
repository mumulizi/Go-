package main

import (
	"fmt"
	"go_dev/day5/homework/model"
)

func main() {
	var (
		name string
		count int
		author string
		createTime string
		search string
		bn string
		bc int
	)

	fmt.Scanln(&name,&count,&author,&createTime)
	cb := model.CreateBook(name,count,author,createTime)

	fmt.Printf("输入查询的书名：")
	fmt.Scanf("%s",&search)
	sb := cb.Search(search)
	fmt.Println(sb)

	fmt.Printf("输入要借的书名和数量：")
	fmt.Scanf("%s %d",&bn,&bc)
	bb := cb.Borrow(bn,bc)
	fmt.Println(bb)

}

