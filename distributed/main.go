package main

import (
	"fmt"
	"go-crawler/distributed/config"
	"go-crawler/distributed/persist/client"
	"go-crawler/engine"
	"go-crawler/scheduler"
	"go-crawler/tuishujun/parser"
)

func main() {
	itemChan, err := client.ItemSaver(
		fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}

	e.Run(engine.Request{
		Url:        "/tags/重生",
		ParserFunc: parser.ParseTag,
	})
}
