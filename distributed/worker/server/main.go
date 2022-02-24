package main

import (
	"fmt"
	"go-crawler/distributed/config"
	"go-crawler/distributed/rpcsupport"
	"go-crawler/distributed/worker"
	"log"
)

func main() {
	log.Fatal(rpcsupport.ServeRpc(
		fmt.Sprintf(":%d", config.WorkerPort0),
		worker.CrawlService{}))
}
