package main

import (
	"golang/crawler/engine"
	"golang/crawler/model"
	"golang/crawler_distributed/config"
	"golang/crawler_distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T){
	const host = ":1234"
	// start ItemSaverServer
	go serveRpc(host, "test1")
	time.Sleep(time.Second)
	// start ItemSaverClient
	client, err := rpcsupport.NewClient(host)
	if err != nil{
		panic(err)
	}
	// Call save
	item := engine.Item{
		Url: "https://album.zhenai.com/u/1734370950",
		Type: "zhenai",
		Id: "1734370950",
		Payload: model.Profile{
			Name: "宁缺毋滥",
			Age: "44",
			Height: "163",
			Education: "大学本科",
			Place: "新疆阿克苏",
			Marriage: "离异",
			Income: "3001-5000元",
		},
	}

	result := ""
	client.Call(config.ItemSaverRPC, item, &result)

	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s",
			result, err)
	}
}
