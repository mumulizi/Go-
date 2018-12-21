package main

import (
	"fmt"
	"sync"
	"time"
)

type task struct {
	n int
}
var (
	m = make(map[int]uint64)
	lock sync.RWMutex
)
func calc(t *task)  {
	var sum uint64
	sum = 1
	for i:= 1; i<=t.n;i++{
		sum *= uint64(i)
	}
	lock.Lock()
	m[t.n] = sum
	lock.Unlock()
}

func main() {
	for i :=0; i<16; i++{
		t := &task{
			n:i,
		}
		go calc(t)
	}
	time.Sleep(5 * time.Second)
	lock.Lock()
	for k,v := range m{
		fmt.Printf("%d! = %v\n",k,v)
	}
	lock.Unlock()
}