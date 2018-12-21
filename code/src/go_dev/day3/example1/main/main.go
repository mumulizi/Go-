package main

import (
	"fmt"
	"strings"
	"time"
)

func urlProcess(url string) string{
	ret := strings.HasPrefix(url,"http://")
	if !ret{
		url = fmt.Sprintf("http://%s",url)
	}
	return url
}
func pathProcess(path string) string {
	ret := strings.HasSuffix(path,"/")
	if !ret{
		path = fmt.Sprintf("%s/",path)
	}
	return path
}

func main()  {
    var (
	    url string
	    path string
    )
	fmt.Scanf("%s%s",&url ,&path)
	url = urlProcess(url)
	fmt.Println(url)
	path =  pathProcess(path)
	fmt.Println(path)
	now := time.Now()
	fmt.Println(now.Format("2006/1/02 15:04"))
	tt :=strings.Trim("abcdadbabdeba", "ab")
	fmt.Println(tt)
}