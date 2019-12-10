package main

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/elastic/go-elasticsearch/v7/esutil"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	clientConfig := elasticsearch.Config{
		Addresses: []string{"http://172.16.4.4:9200"},
	}
	es, err := elasticsearch.NewClient(clientConfig)

	if err != nil {
		panic(err)
	}

	var counter uint64
	wg := sync.WaitGroup{}

	ticker := time.NewTicker(time.Second * 5)
	go func() {
		for range ticker.C {
			log.Println(atomic.LoadUint64(&counter))
		}
	}()

	for tread := 0; tread < 4; tread++ {
		go func() {
			wg.Add(1)
			for i := 0; i < 10000; i++ {
				res, err := esapi.IndexRequest{
					Index: "today",
					Body: esutil.NewJSONReader(struct {
						Name string
						Age  int
					}{
						"Farid",
						i,
					}),
					//DocumentID:   ,
					DocumentType: "person",
					//Refresh:      "true",
					Timeout: time.Second * 2,
				}.Do(context.Background(), es)

				if err != nil {
					log.Println(err)
					continue
				}

				if res.StatusCode == 400 {
					log.Println("wrong")
					continue
				}
				res.Body.Close()
				atomic.AddUint64(&counter, 1)
				wg.Done()
			}
		}()
	}

	wg.Wait()

	res, err := esapi.CountRequest{
		Index:        []string{"today"},
		DocumentType: []string{"person"},
	}.Do(context.Background(), es)

	if err != nil {
		panic(err)
	}
	fmt.Println(res.String())
	res.Body.Close()

	//var size int = 10000
	//
	//res, err := esapi.SearchRequest{
	//	Index:        []string{"today"},
	//	DocumentType: []string{"person"},
	//	Size:         &size,
	//}.Do(context.Background(), es)
	//
	//if err != nil {
	//	panic(err)
	//}
	//res.Body.Close()
}
