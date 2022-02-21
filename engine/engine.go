package engine

import (
	"go-crawler/fetcher"
	"log"
)

const baseUrl = "https://www.tuishujun.com"

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("Fetching %s", baseUrl + r.Url)
		body, err := fetcher.Fetch(baseUrl + r.Url)
		if err != nil {
			log.Printf("Fetcher: error Fetching url %s: %v",
				r.Url, err)
			continue
		}

		parserResult := r.ParserFunc(body)
		requests = append(requests, parserResult.Requests...)

		for _, item := range parserResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}
