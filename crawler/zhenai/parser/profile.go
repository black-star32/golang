package parser

import (
	"golang/crawler/engine"
	"golang/crawler/model"
	"regexp"
)

var nameRe = regexp.MustCompile(`<th><a href="http://album.zhenai.com/u/[0-9a-z]+" [^>]*>([^<]+)</a></th>`)
var genderRe = regexp.MustCompile(`<<td width="180"><span class="grayL">性别：[^>]*>([^<]+)</td>`)
var placeRe = regexp.MustCompile(`<td><span class="grayL">居住地：[^>]*>([^<]+)</td>`)
var ageRe = regexp.MustCompile(`<tr><td width="180"><span class="grayL">年龄：[^>]*>([^<]+)</td>`)
var eduRe = regexp.MustCompile(`<td><span class="grayL">学.*历：[^>]*>([^<]+)</td>`)
var marrRe = regexp.MustCompile(`<td width="180"><span class="grayL">婚况：[^>]*>([^<]+)</td>`)
var heightRe = regexp.MustCompile(`<td width="180"><span class="grayL">身.*高：[^>]*>([^<]+)</td>`)
var incomeRe = regexp.MustCompile(`<td><span class="grayL">月.*薪：[^>]*>([^<]+)</td>`)

func ParseProfile(contents []byte) engine.ParseResult{
	profile := model.Profile{}
	

	profile.Age = extractString(contents, ageRe)
	profile.Height = extractString(contents, heightRe)
	profile.Name = extractString(contents, nameRe)
	profile.Gender = extractString(contents, genderRe)
	profile.Education = extractString(contents, eduRe)
	profile.Place = extractString(contents, placeRe)
	profile.Marriage = extractString(contents, marrRe)
	profile.Income = extractString(contents, incomeRe)
	
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string{
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	}else{
		return ""
	}
}
