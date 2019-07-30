package main

import (
	"fmt"
	"learnGo/retriever/mock"
	real2 "learnGo/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

//接口是隐式的
func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"this is a fake immoc.com"}
	r = real2.Retriever{}
	fmt.Println(download(r))
}
