package engine

import (
	"code/CrawlerConcurrent/fetcher"
	"log"
)

type ConcurrentSimpleEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type SimpleScheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
	//WorkerReady(chan Request)
	//Run()
}

func (e *ConcurrentSimpleEngine) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParserResult)

	// 配置调度的主 channel

	// 创建 worker 协程
	for i := 0; i < e.WorkerCount; i++ {
		createWorkerSimple(in, out)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	for {
		// 下面 in <- 超过 workerCount 数量就会阻塞，导致 <-out chan 中没法继续写入数据，进而导致 <-out无法获取数据进阻塞
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item #%d: %v\n", itemCount, item)
			itemCount++
		}
		// 批量写入 scheduler workerChan ，返回的 request 数量超过 workerCount 就会形成阻塞，无法继续写入
		// 需要开启协程 异步写入 scheduler workerChan，就不会阻塞 ，也不影响 <-out chan的输出
		for _, r := range result.Requests {
			// 因为 worker 执行慢导致，超过 workerCount 的 in<- chan 无法写入而阻塞后续代码，阻塞后续 <-out 获取数据
			e.Scheduler.Submit(r)
		}
	}
}

func createWorkerSimple(in chan Request, out chan ParserResult) {
	go func() {
		for {
			request := <-in
			result, err := workerSimple(request)
			if err != nil {
				continue
			}
			// 将结果送给 engine，如果 request 超过 workerCount 数量， <-out 被阻塞无法获取数据，导致 out<- 无法写入数据而阻塞，形成循环阻塞
			out <- result
		}
	}()
}

func workerSimple(r Request) (ParserResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher error: fetch url=%s, err=%v", r.Url, err)
		return ParserResult{}, err
	}

	return r.ParserFunc(body), nil
}
