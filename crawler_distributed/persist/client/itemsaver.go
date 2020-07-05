package client

import (
	"golang/crawler/engine"
	"golang/crawler_distributed/config"
	"golang/crawler_distributed/rpcsupport"
	"log"
)

func ItemSaver(host string) (chan engine.Item, error){
	clinet, err := rpcsupport.NewClient(host)
	if err != nil{
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <- out
			log.Printf("Item Saver: got item " +
				"#%d: %v", itemCount, item)
			itemCount++

			// Call RPC to save item
			result := ""
			err := clinet.Call(config.ItemSaverRPC, item, &result)
			if err != nil {
				log.Printf("Item Saver: error " +
					"saving item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}
