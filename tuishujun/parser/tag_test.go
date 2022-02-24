package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseTag(t *testing.T) {
	contents, err := ioutil.ReadFile("tag_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseTag(contents, "")

	if len(result.Requests) != 15 {
		t.Errorf("Requests should contain 15 "+
			"request; but was %v", len(result.Requests))
	}
}
