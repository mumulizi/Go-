package model

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
	if n == p.Name{
		ret := fmt.Sprintf("当前图书:%s,数量：%d,作者:%s,出版日期:%s\n",p.Name,p.Count,p.Author,p.CreateTime)
		return ret
	}
	return "NOT EXITS"
}

func (p *Book) Borrow(n string,c int) string  {
	p.Count = p.Count -c
	ret := fmt.Sprintf("当前图书:%s,数量：%d,作者:%s,出版日期:%s\n",p.Name,p.Count,p.Author,p.CreateTime)
	return ret
}
