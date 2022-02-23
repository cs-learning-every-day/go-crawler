package parser

import (
	"go-crawler/engine"
	"go-crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseBook(t *testing.T) {
	contents, err := ioutil.ReadFile("book_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseBook(contents, "https://www.tuishujun.com/books/989")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 "+
			"element; but was %v", result.Items)
	}

	actual := result.Items[0]

	expected := engine.Item{
		Url:  "https://www.tuishujun.com/books/989",
		Type: "tuishujun",
		Id:   "989",
		Payload: model.Book{
			Title:        "仙逆",
			Author:       "耳根",
			Grade:        "8.4",
			WordCount:    "652 万字",
			ChapterCount: "2072章\r",
			State:        "完结",
			UpdateTime:   "11 年前\r",
			Source:       "起点中文网",
			Intro:        "顺为凡，逆则仙，只在心中一念间……请看耳根作品",
			CoverImgUrl:  "https://tuishujun.s3.ladydaily.com/cover/989?fmt=webp&amp;w=90",
		},
	}

	if actual != expected {
		t.Errorf("expected %v; but was %v",
			expected, actual)
	}
}
