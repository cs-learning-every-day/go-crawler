package main

import (
	"go-crawler/engine"
	"go-crawler/scheduler"
	"go-crawler/tuishujun/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}

	e.Run(engine.Request{
		Url:        "/tags/言情?page=6",
		ParserFunc: parser.ParseTag,
	})
}
