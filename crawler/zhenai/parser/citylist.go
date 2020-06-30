package parser

import (
	"golang/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult{
	re := regexp.MustCompile(cityListRe)

	result := engine.ParseResult{}
	mathces := re.FindAllSubmatch(contents, -1)
	for _, m := range mathces{
		result.Items = append(
			result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParseFunc: engine.NilParse,
		})
	}
	return result
}
