package main

import "testing"
//Testxxx 必须这么写，参数也必须这么写，运行的时候直接go test
func TestAdd(t *testing.T) {
	r := Add(2,4)
	if r != 6 {
		t.Fatalf("add(2,4)err,expect:%d,actual:%d",6,r)
	}
	t.Logf("test add success")
}