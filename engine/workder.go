package engine

import (
	"go-crawler/fetcher"
	"log"
)

const baseUrl = "https://www.tuishujun.com"

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", baseUrl+r.Url)
	body, err := fetcher.Fetch(baseUrl + r.Url)
	if err != nil {
		log.Printf("Fetcher: error Fetching url %s: %v",
			r.Url, err)
		return ParseResult{}, err
	}

	return r.ParserFunc(body), nil
}