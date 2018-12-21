package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main()  {
	conn ,err := net.Dial("tcp","localhost:7000")
	if err != nil{
		fmt.Println("Error dialing.",err.Error())
		return
	}
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input , _ := inputReader.ReadString('\n')
		timmedInput := strings.Trim(input,"\r\n")
		if timmedInput == "Q"{
			return
		}
		_,err := conn.Write([]byte(timmedInput))
		if err != nil{
			return
		}
	}


}

