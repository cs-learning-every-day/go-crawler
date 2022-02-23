package engine

import (
	"go-crawler/config"
	"go-crawler/fetcher"
	"log"
)

func worker(r Request) (ParseResult, error) {
	body, err := fetcher.Fetch(config.BaseUrl + r.Url)
	if err != nil {
		log.Printf("Fetcher: error Fetching url %s: %v",
			r.Url, err)
		return ParseResult{}, err
	}

	return r.ParserFunc(body), nil
}
