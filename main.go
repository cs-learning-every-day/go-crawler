package main

import (
	"go-crawler/engine"
	"go-crawler/tuishujun/parser"
)

func main() {
	engine.Run(engine.Request{
		Url: "/hot-tags",
		ParserFunc: parser.ParseTagList,
	})
}
