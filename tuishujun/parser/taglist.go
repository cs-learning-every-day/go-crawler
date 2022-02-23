package parser

import (
	"go-crawler/engine"
	"regexp"
)

const tagListRe = `<a href="(/tags/.+)" title="好看的(.+)小说推荐_排行" class="tag-group__tags__tag"`

func ParseTagList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(tagListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	limit := 0
	for _, m := range matches {
		if limit > 10 {
			break
		}
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests,
			engine.Request{
				Url:        string(m[1]),
				ParserFunc: ParseTag,
			})
		//fmt.Printf("Tag: %s\tURL: %s\n", m[2], m[1])
		limit += 1
	}
	//fmt.Printf("Matches found: %d\n", len(matches))
	return result
}
