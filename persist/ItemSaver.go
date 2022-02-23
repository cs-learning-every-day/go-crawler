package persist

import (
	"context"
	"errors"
	"go-crawler/engine"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSaver() chan engine.Item {
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			itemCount++
			item := <-out
			log.Printf("Item Saver: got item #%d: %v\n", itemCount, item)

			err := save(item)
			if err != nil {
				log.Printf("Item Saver: error saving item %v: %v\n", item, err)
			}
		}
	}()
	return out
}

func save(item engine.Item) error {
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.10.103:9200/"),
		// ust turn off sniff in docker
		elastic.SetSniff(false),
	)
	if err != nil {
		return err
	}

	if item.Type == "" {
		return errors.New("must supply Type")
	}

	indexService := client.Index().
		Index("xiaoshuo_data").
		Type(item.Type).
		BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err = indexService.Do(context.Background())

	if err != nil {
		return err
	}

	return nil
}
