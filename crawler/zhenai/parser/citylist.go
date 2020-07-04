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
		//result.Items = append(
		//	result.Items, "City " + string(m[2]))
		url := string(m[1])
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParseFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, url)
			},
		})
	}
	return result
}
