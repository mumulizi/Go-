package main

import (
	"fmt"
	"go_dev/structAtm/utils"
)

func main() {
	fmt.Println("面向对象ATM练习")
	utils.NewStructAtm().MainMenu()
}