package persist

import (
	"context"
	"encoding/json"
	"golang/crawler/model"
	"gopkg.in/olivere/elastic.v5"
	"testing"
)

func TestSaver(t *testing.T) {
	profile := model.Profile{
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
	}
	id, err := save(profile)

	if err != nil {
		panic(err)
	}

	client, err := elastic.NewClient(
		elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	resp, err := client.Get().
		Index("dating_profile").
		Type("zhenai").Id(id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	t.Logf("%s", resp.Source)

	var actual model.Profile
	err = json.Unmarshal(
		*resp.Source, &actual)

	if err != nil{
		panic(err)
	}

	if actual != profile {
		t.Errorf("got %v; expected %v", actual, profile)
	}
}
