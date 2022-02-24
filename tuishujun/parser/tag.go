package parser

import (
	"go-crawler/engine"
	"regexp"
)

var tagRe = regexp.MustCompile(`<a href="(/books/\d*)" class="tsj-book-item__info__right__book-name font-normal no-underline text-color-high cursor-pointer block" .*<h2>\s*(\S+)\s*</h2></a>`)

func ParseTag(contents []byte, _ string) engine.ParseResult {
	matches := tagRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		url := string(m[1])
		result.Requests = append(result.Requests,
			engine.Request{
				Url: url,
				Parser: engine.NewFuncParser(
					ParseBook, "ParseBook"),
			})
	}
	return result
}
