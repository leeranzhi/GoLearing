package main

import (
	"fmt"
	"learnGo/retriever/mock"
	real2 "learnGo/retriever/real"
	"time"
)

//接口变量自带指针
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
	inspect(r)

	r = &real2.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)

	if realRetriever, ok := r.(*real2.Retriever); ok {
		fmt.Println(realRetriever.TimeOut)
	} else {
		fmt.Println("not a mock retriever")
	}

	//fmt.Println(download(r))
}

func inspect(retriever Retriever) {

	fmt.Printf("%T %v\n", retriever, retriever)

	switch v := retriever.(type) {
	case mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real2.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}

}
