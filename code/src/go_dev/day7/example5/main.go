package main

import (
	"flag"
	"fmt"
)

func main() {
	var Name string
	var Port int

	flag.StringVar(&Name, "n","","Please your start service name.")
	flag.IntVar(&Port,"d",1000,"please your port")

	flag.Parse()
	fmt.Println("Name",Name)
	fmt.Println("Port",Port)
}


