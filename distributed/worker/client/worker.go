package client

import (
	"go-crawler/distributed/config"
	"go-crawler/distributed/worker"
	"go-crawler/engine"
	"net/rpc"
)

func CreateProcessor(clients chan *rpc.Client) engine.Processor {
	return func(req engine.Request) (
		engine.ParseResult, error) {

		sReq := worker.SerializeRequest(req)

		var sResult worker.ParseResult
		c := <-clients
		err := c.Call(config.CrawlServiceRpc, sReq, &sResult)

		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult), nil
	}
}
