package main

import (
	"flag"
	"fmt"
	"go-crawler/config"
	config2 "go-crawler/distributed/config"
	"go-crawler/distributed/persist"
	"go-crawler/distributed/rpcsupport"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

var port = flag.Int("port", 0,
	"the port for me to listen on")

func main() {
/*	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}*/
	log.Fatal(serveRpc(
		fmt.Sprintf(":%d", config2.ItemSaverPort),
		config.ElasticIndex))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetURL(config.DockerUrl),
		// ust turn off sniff in docker
		elastic.SetSniff(false),
	)
	if err != nil {
		return err
	}

	return rpcsupport.ServeRpc(host,
		&persist.ItemSaverService{
			Client: client,
			Index:  index,
		})
}
