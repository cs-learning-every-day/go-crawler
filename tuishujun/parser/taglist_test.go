package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseTagList(t *testing.T) {
	contents, _ := ioutil.ReadFile("taglist_test_data.html")

	result := ParseTagList(contents)
	const resultSize = 1411
	if len(result.Requests) != resultSize {
		t.Errorf("expected result size %d; but had %d", resultSize, len(result.Requests))
	}

	expectedUrls := []string{
		"/tags/重生", "/tags/都市", "/tags/轻松",
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but was %s",
				i, url, result.Requests[i].Url)
		}
	}
}
