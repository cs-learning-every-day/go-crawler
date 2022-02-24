package main

import (
	"fmt"
	config2 "go-crawler/distributed/config"
	"go-crawler/distributed/rpcsupport"
	"go-crawler/distributed/worker"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	// TODO: Use a fake fetcher to handle the url.
	// So we don't get data from tuishujun
	req := worker.Request{
		Url: "/books/989",
		Parser: worker.SerializedParser{
			Name: config2.ParseBook,
			Args: "",
		},
	}
	var result worker.ParseResult
	err = client.Call(config2.CrawlServiceRpc, req, &result)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}

	// TODO: Verify results
}
