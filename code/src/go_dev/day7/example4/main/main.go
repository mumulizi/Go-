package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	inputFile := "C:/a"
	OutputFile := "C:/b"

	buf,err := ioutil.ReadFile(inputFile)
	if err != nil{
		fmt.Fprintf(os.Stderr,"File Error%s\n",err)
		return
	}


	fmt.Printf("%s \n",string(buf))
	err = ioutil.WriteFile(OutputFile,buf,0x1)
	if err != nil{
		panic(err.Error())
		//return
	}

}
