package main

import (
	"fmt"
	"go-crawler/distributed/config"
	itemsaver "go-crawler/distributed/persist/client"
	worker "go-crawler/distributed/worker/client"
	"go-crawler/engine"
	"go-crawler/scheduler"
	"go-crawler/tuishujun/parser"
)

func main() {
	itemChan, err := itemsaver.ItemSaver(
		fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	processor := worker.CreateProcessor()

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}

	e.Run(engine.Request{
		Url:    "/tags/重生",
		Parser: engine.NewFuncParser(parser.ParseTag, "ParseTag"),
	})
}
