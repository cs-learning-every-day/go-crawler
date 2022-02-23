package model

import "encoding/json"

type Book struct {
	Title        string
	Author       string
	CoverImgUrl  string
	Grade        string
	WordCount    string
	ChapterCount string
	State        string
	UpdateTime   string
	Source       string // 来源
	Intro        string
}

func FromJsonObj(o interface{}) (Book, error) {
	var book Book
	s, err := json.Marshal(o)
	if err != nil {
		return book, err
	}
	err = json.Unmarshal(s, &book)
	return book, err
}
