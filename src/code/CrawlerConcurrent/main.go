package main

import (
	"code/CrawlerConcurrent/engine"
	"code/CrawlerConcurrent/persist"
	"code/CrawlerConcurrent/scheduler"
	"code/CrawlerConcurrent/website/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    persist.ItemSaver(),
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
