package engine

import (
	"code/CrawlerConcurrent/fetcher"
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan interface{}
}

type Scheduler interface {
	ReadNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParserResult)
	// 启动 scheduler，等待 request 输入并分配给 workerChan 等待 worker 执行
	e.Scheduler.Run()

	// 创建 worker 协程
	for i := 0; i < e.WorkerCount; i++ {
		// 每个 worker 有自己的 request chan 调度
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			//log.Printf("EngineOuter get item #: %v\n", item)
			//go func() {
				e.ItemChan <- item
				//log.Printf("EngineInner get item #: %v\n", item)
			//}()
		}
		for _, r := range result.Requests {
			e.Scheduler.Submit(r)
		}
	}
}

func createWorker(in chan Request, out chan ParserResult, ready ReadNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

func worker(r Request) (ParserResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher error: fetch url=%s, err=%v", r.Url, err)
		return ParserResult{}, err
	}

	return r.ParserFunc(body), nil
}
