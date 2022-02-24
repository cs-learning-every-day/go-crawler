package view

import (
	"go-crawler/engine"
	"go-crawler/frontend/model"
	model2 "go-crawler/model"
	"os"
	"testing"
)

func TestSearchResultView_Render(t *testing.T) {
	view := CreateSearchResultView(
		"template.html")

	out, err := os.Create("template.test.html")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	page := model.SearchResult{}
	page.Hits = 123
	item := engine.Item{
		Url:  "https://www.tuishujun.com/books/989",
		Type: "tuishujun",
		Id:   "989",
		Payload: model2.Book{
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
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	err = view.Render(out, page)
	if err != nil {
		t.Error(err)
	}
}
