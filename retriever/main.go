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

type Poster interface {
	Post(url string,
		form map[string]string) string
}

const url = "http://www.imooc.com"

func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name":   "ccmouse",
			"course": "golang",
		})
}

//接口的组合
type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}

//接口是隐式的
func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retriever
	retriever := mock.Retriever{"this is a fake immoc.com"}
	r = &retriever
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

	fmt.Println("Try a seesion")
	fmt.Println(session(&retriever))

	//fmt.Println(download(r))
}

func inspect(retriever Retriever) {

	fmt.Printf("%T %v\n", retriever, retriever)

	switch v := retriever.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real2.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}

}
