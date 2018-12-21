package main

import "fmt"

type Book struct {
	Name string
	Count int
	Author string
	CreateTime string
}

func CreateBook(name string,count int,author string,createTime string) (b *Book) {
	b = &Book{
		Name:name,
		Count:count,
		Author:author,
		CreateTime:createTime,
	}
	//fmt.Println(b.Name)

	return
}

func (p *Book) Search(n string) string  {
	fmt.Println("n:",n)
	fmt.Println("p:",p.Name)
	//if n == p.Name{
	//	ret := fmt.Sprintf("当前图书:%s,数量：%d,作者:%s,出版日期:%s\n",p.Name,p.Count,p.Author,p.CreateTime)
	//	return ret
	//}
	return "NOT EXITS"
}




func (p *Book) Borrow(n string,c int) string  {
	fmt.Println("--------")
	return "1111"
}


func main() {
	var (
		name string
		count int
		author string
		createTime string
		//b Book
		search string
		b string
	)

	fmt.Scanln(&name,&count,&author,&createTime)
	ret := CreateBook(name,count,author,createTime)
	//fmt.Println("b:",r)
	fmt.Println("输入查询的书名：")
	fmt.Scanf("%s",&search)
	//ret := b.Search(&search)
	//fmt.Println(ret)


	//ret := CreateBook("go",11,"alex","2010-10-10")
	//var b Book
	//b.Name = "go"
	//b.Count=10
	//b.Author="alex"
	//b.CreateTime="2019-01-10"

	fmt.Println("newB:",ret)
	r :=ret.Search(search)
	fmt.Println(r)



	fmt.Println("输入c的书名：")
	fmt.Scanf("%s",&b)
}


