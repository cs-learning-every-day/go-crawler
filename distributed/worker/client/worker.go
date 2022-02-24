package client

import (
	"fmt"
	"go-crawler/distributed/config"
	"go-crawler/distributed/rpcsupport"
	"go-crawler/distributed/worker"
	"go-crawler/engine"
)

func CreateProcessor() engine.Processor {
	client, err := rpcsupport.NewClient(
		fmt.Sprintf(":%d", config.WorkerPort0))
	if err != nil {
		panic(err)
	}

	return func(req engine.Request) (
		engine.ParseResult, error) {

		sReq := worker.SerializeRequest(req)

		var sResult worker.ParseResult

		err := client.Call(config.CrawlServiceRpc, sReq, &sResult)

		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult), nil
	}
}
