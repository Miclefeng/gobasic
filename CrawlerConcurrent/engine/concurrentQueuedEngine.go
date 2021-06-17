package engine

import (
	"code/CrawlerConcurrent/fetcher"
	"log"
)

type ConcurrentQueuedEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type QueuedScheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
	WorkerReady(chan Request)
	Run()
}

func (e *ConcurrentQueuedEngine) Run(seeds ...Request) {
	out := make(chan ParserResult)

	// 启动 scheduler，等待 request 输入并分配给 workerChan 等待 worker 执行
	e.Scheduler.Run()

	// 创建 worker 协程
	for i := 0; i < e.WorkerCount; i++ {
		createWorkerQueued(out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item #%d: %v\n", itemCount, item)
			itemCount++
		}
		for _, r := range result.Requests {
			e.Scheduler.Submit(r)
		}
	}
}

func createWorkerQueued(out chan ParserResult, s Scheduler) {
	in := make(chan Request)
	go func() {
		for {
			s.WorkerReady(in)
			request := <-in
			result, err := workerQueued(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

func workerQueued(r Request) (ParserResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher error: fetch url=%s, err=%v", r.Url, err)
		return ParserResult{}, err
	}

	return r.ParserFunc(body), nil
}
