package main

import (
	"code/CrawlerSingle/engine"
	"code/CrawlerSingle/website/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
