package parser

import (
	"code/CrawlerSingle/fetcher"
	"fmt"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", contents)
	ParseCityList(contents)
}
