package main

import (
	"fmt"
	"golang/crawler_distributed/config"
	"golang/crawler_distributed/persist"
	"golang/crawler_distributed/rpcsupport"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func main() {
	log.Fatal(serveRpc(fmt.Sprintf(":%d", config.ItemSaverPort), config.ElasticIndex))
}

func serveRpc(host, index string) error{
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil{
		return err
	}
	return rpcsupport.ServeRPC(host, &persist.ItemSaverService{
		Client: client,
		Index: index,
	})
}
