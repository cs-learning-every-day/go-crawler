package parser

import (
	"go-crawler/engine"
	"regexp"
)

var tagRe = regexp.MustCompile(`<a href="(/books/\d*)" class="tsj-book-item__info__right__book-name font-normal no-underline text-color-high cursor-pointer block" .*<h2>\s*(\S+)\s*</h2></a>`)

func ParseTag(contents []byte) engine.ParseResult {
	matches := tagRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "Book Name: "+string(m[2]))
		result.Requests = append(result.Requests,
			engine.Request{
				Url:        string(m[1]),
				ParserFunc: ParseBook,
			})
	}
	return result
}
