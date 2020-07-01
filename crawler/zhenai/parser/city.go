package parser

import (
	"golang/crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult{
	re := regexp.MustCompile(cityRe)

	result := engine.ParseResult{}
	mathces := re.FindAllSubmatch(contents, -1)
	for _, m := range mathces{
		result.Items = append(
			result.Items, "User " + string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParseFunc: engine.NilParse,
		})
	}
	return result
}
