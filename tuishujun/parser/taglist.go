package parser

import (
	"go-crawler/engine"
	"regexp"
)

const tagListRe = `<a href="(/tags/.+)" title="好看的(.+)小说推荐_排行" class="tag-group__tags__tag"`

func ParseTagList(contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(tagListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests,
			engine.Request{
				Url:    string(m[1]),
				Parser: engine.NewFuncParser(ParseTag, "ParseTag"),
			})
	}
	return result
}
