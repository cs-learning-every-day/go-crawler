package main

import (
	"go-crawler/engine"
	"go-crawler/tuishujun/parser"
)

func main() {
	engine.Run(engine.Request{
		Url: "/books/104816",
		ParserFunc: parser.ParseBook,
	})
}
