package parser

import (
	"code/CrawlerConcurrent/fetcher"
	"fmt"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun/aba")
	if err != nil {
		panic(err)
	}
	profile := ParseProfile(contents)
	fmt.Println(profile)
}
