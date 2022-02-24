package main

import (
	"fmt"
	"go-crawler/distributed/config"
	"go-crawler/distributed/rpcsupport"
	"go-crawler/engine"
	"go-crawler/model"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	host := fmt.Sprintf(":%d", config.ItemSaverPort)

	// start ItemSaverServer
	go serveRpc(host, "test1")
	time.Sleep(time.Second)

	// start ItemSaverClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	// Call save
	item := engine.Item{
		Url:  "https://www.tuishujun.com/books/989",
		Type: "tuishujun",
		Id:   "989",
		Payload: model.Book{
			Title:        "仙逆",
			Author:       "耳根",
			Grade:        "8.4",
			WordCount:    "652 万字",
			ChapterCount: "2072章",
			State:        "完结",
			UpdateTime:   "11 年前",
			Source:       "起点中文网",
			Intro:        "顺为凡，逆则仙，只在心中一念间……请看耳根作品",
			CoverImgUrl:  "https://tuishujun.s3.ladydaily.com/cover/989?fmt=webp&amp;w=90",
		},
	}

	result := ""
	err = client.Call(config.ItemSaverRpc, item, &result)

	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s",
			result, err)
	}
}
