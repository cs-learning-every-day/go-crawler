package parser

import (
	"go-crawler/engine"
	"regexp"
)

const tagListRe = `<a href="(/tags/.+)" title="好看的(.+)小说推荐_排行" class="tag-group__tags__tag"`

func ParseTagList(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(tagListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}
	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests,
			engine.Request{
				Url:        string(m[1]),
				ParserFunc: engine.NilParser,
			})
		//fmt.Printf("Tag: %s\tURL: %s\n", m[2], m[1])
	}
	//fmt.Printf("Matches found: %d\n", len(matches))
	return result
}
