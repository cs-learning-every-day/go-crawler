package main

import (
	"go-crawler/config"
	"go-crawler/engine"
	"go-crawler/persist"
	"go-crawler/scheduler"
	"go-crawler/tuishujun/parser"
)

func main() {
	itemChan, err := persist.ItemSaver(config.ElasticIndex)
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}

	e.Run(engine.Request{
		Url:    "/hot-tags",
		Parser: engine.NewFuncParser(parser.ParseTagList, "ParseTagList"),
	})
}
