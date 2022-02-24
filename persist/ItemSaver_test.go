package persist

import (
	"context"
	"encoding/json"
	"go-crawler/config"
	"go-crawler/engine"
	"go-crawler/model"
	"gopkg.in/olivere/elastic.v5"
	"testing"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
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

	client, err := elastic.NewClient(
		elastic.SetURL(config.DockerUrl),
		// ust turn off sniff in docker
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}

	const index = "dating_test"

	err = Save(client, index, expected)
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	t.Logf("%s\n", resp.Source)

	var actual engine.Item
	json.Unmarshal(*resp.Source, &actual)

	actualPayload, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualPayload

	// Verify result
	if actual != expected {
		t.Errorf("got %v; expected %v",
			actual, expected)
	}

}
