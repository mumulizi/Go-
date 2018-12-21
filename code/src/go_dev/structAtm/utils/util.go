package utils

import (
	"fmt"
	"strings"
)

type StructAtm struct {
	action int
	loop bool
	save_money int
	money int
	draw_money int
}

func NewStructAtm() *StructAtm  {
	return &StructAtm{
		action : 0,
		loop :true,
		save_money :0,
		money :0,
		draw_money :0,
	}
}


func (this *StructAtm) find(){
	fmt.Printf("余额为：%d \n",this.money)
}
func (this *StructAtm) save(){
	fmt.Println("请输入存款金额：")
	//fmt.Scanf("%d\n",&save_money) //使用scanf必须加\n不然会打印两次
	fmt.Scanln(&this.save_money)
	fmt.Println(this.save_money)
	this.money += this.save_money
	fmt.Printf("存款金额%d，存款成功\n",this.save_money)
}
func (this *StructAtm) draw(){
	fmt.Println("请输入具体取款金额：")
	fmt.Scanln(&this.draw_money)
	if this.draw_money <= this.money{
		this.money -= this.draw_money
		fmt.Printf("成功取款,余额为：%d\n",this.money)
	}else {
		fmt.Println("余额不足")
	}
}
func (this *StructAtm) shoping(){
	shop_goods := make(map[string]int)
	goods_list := make([]string,0)
	var goodName string
	var shopmany int
	if len(shop_goods)==0 {
		this.add_goods(&shop_goods)
	}
	fmt.Println(shop_goods)
	this.shop_car(goodName ,&goods_list)
	for i:=0 ; i<len(goods_list);i++{
		v,_ :=shop_goods[goods_list[i]]
		shopmany += v
	}
	if this.money < shopmany{
		fmt.Println("购物失败，余额不足，请及时充值")
	}else{
		this.money -=  shopmany
	}
}

func (this *StructAtm) add_goods(g *map[string]int) {
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
func (this *StructAtm) shop_car(gn string,gl *[]string) {
	for{
		fmt.Println("输入要购买的商品(Q退出)：")
		fmt.Scanln(&gn)
		if gn =="Q" {
			break
		}
		*gl = append(*gl,gn)
	}
}
func (this *StructAtm) manager() {
	fmt.Println("-----欢迎登陆中国银行-----")
	fmt.Println("\t1:查询余额")
	fmt.Println("\t2:存款")
	fmt.Println("\t3:取款")
	fmt.Println("\t4:购物")
	fmt.Println("\t5:退出")
	fmt.Println("请输入你的操作：")
	fmt.Scanln(&this.action)
	switch this.action {
	case 1:
		this.find()
	case 2:
		this.save()
	case 3:
		this.draw()
	case 4:
		this.shoping()
	case 5:
		this.loop = false
	default:
		fmt.Println("输入有误，请输入1-4之间的数字")
	}
}

func (this *StructAtm) MainMenu() {
	for {
		this.manager()
		if this.loop == false{
			fmt.Println("----欢迎使用，再见----")
			break
		}
	}
}