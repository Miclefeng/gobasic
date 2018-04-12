package main

import (
	"fmt"
	"miclefeng/learngo/retriever/mock"
	"miclefeng/learngo/retriever/real"
	"time"
)
// interface 里面有实现者类型和实现者(值/指针)
// 接口变量自带指针
// 接口变量同样采用值传递，几乎不需要使用接口的指针
// 指针接收者实现只能以指针方式使用；值接收者都可以
type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

type RetrieverPoster interface {
	Retriever
	Poster
}

const url  = "https://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster)  {
	poster.Post(url, map[string]string{
		"name" : "miclefeng",
		"course" : "golang"})
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string {
		"contents" : "anther faked imooc.com",
	})

	return s.Get(url)
}

func main() {
	var r Retriever
	retriever := mock.Retriever{"this is a fake imooc.com"}
	r = &retriever
	inspect(r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimOut: time.Minute,
	}
	inspect(r)
	fmt.Println()

	// Type assertion
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("Not a Mock Retriever")
	}

	//realRetriever := r.(*real.Retriever)
	//fmt.Println(realRetriever.TimOut)
	//fmt.Println(download(r))
	fmt.Println("Try a session :")
	fmt.Println(session(&retriever))
}

func inspect(r Retriever) {
	fmt.Println("Inspecting: ", r)
	fmt.Printf(" > %T %v\n", r, r)
	fmt.Print(" > Type switch: ")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents: ", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent: ", v.UserAgent)
	}
	fmt.Println()
}
