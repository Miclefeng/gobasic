package parser

import (
	"code/CrawlerConcurrent/engine"
	"regexp"
)

const cityListReg = `<a[^>]+href="(http://www.zhenai.com/zhenghun/[\w]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParserResult {
	compile := regexp.MustCompile(cityListReg)
	matches := compile.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}
	//limiter := 10
	for _, m := range matches {
		// 城市名称添加到元素列表
		//result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseProfile,
		})
		//if limiter < 0 {
		//	break
		//}
		//limiter--
		//fmt.Printf("City: %s, URL: %s\n", m[2], m[1])
	}
	//fmt.Printf("Matches found %d\n", len(matches))
	return result
}
