package main

import (
	"go-crawler/engine"
	"go-crawler/persist"
	"go-crawler/scheduler"
	"go-crawler/tuishujun/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan: persist.ItemSaver(),
	}

	e.Run(engine.Request{
		Url:        "/hot-tags",
		ParserFunc: parser.ParseTagList,
	})
}
