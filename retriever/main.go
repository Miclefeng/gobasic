package main

import (
	"fmt"
	"miclefeng/learngo/retriever/mock"
	"miclefeng/learngo/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("https://www.imooc.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"this is a fake imooc.com"}
	r = real.Retriever{}
	fmt.Println(download(r))
}
