package worker

import (
	"golang/crawler/engine"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

type CrawlService struct {

}

type Request struct {
	Url    string
	Parser SerializedParser
}

type ParseResult struct {
	Items []engine.Item
	Requests []Request
}

func SerializeRequest(r engine.Request) Request{
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}

	for _, req := range r.Requests{
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

func DeserializeRequest(r Request) engine.Request{
	return engine.Request{
		Url: r.Url,
		Parser: deserializeParser(r.Parser),
	}
}

func DeserializeResult(r ParseResult) engine.ParseResult{
	result := engine.ParseResult{
		Items: r.Items,
	}

	for _, req := range r.Requests{
		result.Requests = append(result.Requests, DeserializeRequest(req))
	}
	return result
}

func deserializeParser(p SerializedParser) engine.Parser {
	switch p.Name {
	//case config.ParseCityList:
	//	return engine.NewFuncParser(parser.ParseCityList, config.ParseCityList), nil
	//default:
	//	return nil, error.New("unknow parser name")
	}
}

func (CrawlService) Process(req engine.Request, result *engine.ParseResult) error{

}