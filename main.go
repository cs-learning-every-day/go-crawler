package main

import (
	"go-crawler/engine"
	"go-crawler/scheduler"
	"go-crawler/tuishujun/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}

	e.Run(engine.Request{
		Url:        "/tags/都市",
		ParserFunc: parser.ParseTag,
	})
}
