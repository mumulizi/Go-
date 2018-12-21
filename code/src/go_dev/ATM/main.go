package main

import (
	"fmt"
	"strings"
)

var (
	action int
	loop bool = true
	save_money int
	money int
	draw_money int
)

func find() {
	fmt.Printf("余额为：%d \n",money)
}
func save() {
	fmt.Println("请输入存款金额：")
	//fmt.Scanf("%d\n",&save_money) //使用scanf必须加\n不然会打印两次
	fmt.Scanln(&save_money)
	fmt.Println(save_money)
	money += save_money
	fmt.Printf("存款金额%d，存款成功\n",save_money)
}
func draw() {
	fmt.Println("请输入具体取款金额：")
	fmt.Scanln(&draw_money)
	if draw_money <= money{
		money -= draw_money
		fmt.Printf("成功取款,余额为：%d\n",money)
	}else {
		fmt.Println("余额不足")
	}
}
func add_goods(g *map[string]int) {
	for{
		var(
			name string
			price int
		)
		fmt.Println("输入新增的商品名称,按Q退出：")
		fmt.Scanln(&name)
		if strings.ToUpper(name)=="Q"{
			break
		}
		fmt.Println("输入该商品价格：")
		fmt.Scanln(&price)
		(*g)[name]= price
	}
}

func shop_car(gn string,gl *[]string)  {
	for{
		fmt.Println("输入要购买的商品(Q退出)：")
		fmt.Scanln(&gn)
		if gn =="Q" {
			break
		}
		*gl = append(*gl,gn)
	}

}
func shoping() {
	shop_goods := make(map[string]int)
	goods_list := make([]string,0)
	var goodName string
	var shopmany int
	if len(shop_goods)==0 {
		add_goods(&shop_goods)
	}
	fmt.Println(shop_goods)
	shop_car(goodName ,&goods_list)
	for i:=0 ; i<len(goods_list);i++{
		v,_ :=shop_goods[goods_list[i]]
		shopmany += v
	}
	if money < shopmany{
		fmt.Println("购物失败，余额不足，请及时充值")
	}else{
		money -=  shopmany
	}
}
func manager()  {
	fmt.Println("-----欢迎登陆中国银行-----")
	fmt.Println("\t1:查询余额")
	fmt.Println("\t2:存款")
	fmt.Println("\t3:取款")
	fmt.Println("\t4:购物")
	fmt.Println("\t5:退出")
	fmt.Println("请输入你的操作：")
	fmt.Scanln(&action)
	switch action {
	case 1:
		find()
	case 2:
		save()
	case 3:
		draw()
	case 4:
		shoping()
	case 5:
		loop = false
	default:
		fmt.Println("输入有误，请输入1-4之间的数字")
	}
}

func main() {
	for {
		manager()
		if loop == false{
			fmt.Println("----欢迎使用，再见----")
			break
		}
	}
}
