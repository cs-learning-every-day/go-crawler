package engine

import (
	"log"
	"strconv"
	"strings"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigureMasterWorkerChan(in)

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}

	for _, seed := range seeds {
		e.Scheduler.Submit(seed)
	}

	count := 0
	for {
		result := <-out

		for _, item := range result.Items {
			log.Printf("Got item #%d: %v\n", count, item)
			count++
		}

		for _, request := range result.Requests {
			// 默认tags页面最多20分页
			if strings.Contains(request.Url, "/tags") {
				oldUrl := request.Url
				for i := 1; i <= 20; i++ {
					request.Url = oldUrl + "?page=" + strconv.Itoa(i)
					e.Scheduler.Submit(request)
				}
			} else {
				e.Scheduler.Submit(request)
			}
		}
	}
}

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
