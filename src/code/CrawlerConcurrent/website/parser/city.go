package parser

import (
	"code/CrawlerConcurrent/engine"
	"regexp"
)

const cityReg = `<th>[^<]*<a href="(http[s]?://album.zhenai.com/u/[\d]+)"[^>]*>([^<]+)</a>[^<]*</th>`

func ParseCity(contents []byte) engine.ParserResult {
	compile := regexp.MustCompile(cityReg)
	matches := compile.FindAllSubmatch(contents, -1)
	var result engine.ParserResult
	for _, m := range matches {
		name := string(m[2])
		result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParser,
		})
	}
	return result
}
