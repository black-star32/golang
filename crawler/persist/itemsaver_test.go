package persist

import (
	"context"
	"encoding/json"
	"golang/crawler/engine"
	"golang/crawler/model"
	"gopkg.in/olivere/elastic.v5"
	"testing"
)

func TestSaver(t *testing.T) {
	expected := engine.Item{
		Url: "https://album.zhenai.com/u/1734370950",
		Type: "zhenai",
		Id: "1734370950",
		Payload: model.Profile{
			//Name string
			//Url string
			//Gender string
			//Age string
			//Height string
			//Education string
			//Place string
			//Marriage string
			//Income string
			Name: "宁缺毋滥",
			Age: "44",
			Height: "163",
			Education: "大学本科",
			Place: "新疆阿克苏",
			Marriage: "离异",
			Income: "3001-5000元",
		},
	}


	client, err := elastic.NewClient(
		elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	const index ="dating_test"
	// save
	err = Save(client, index, expected)

	if err != nil {
		panic(err)
	}


	// fetch
	resp, err := client.Get().
		Index(index).
		Type(expected.Type).Id(expected.Id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	t.Logf("%s", resp.Source)

	var actual engine.Item
	err = json.Unmarshal(
		*resp.Source, &actual)

	if err != nil{
		panic(err)
	}

	actualProfile, _ := model.FromJsonObj(actual.Payload)

	actual.Payload = actualProfile

	// verify result
	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}
}
