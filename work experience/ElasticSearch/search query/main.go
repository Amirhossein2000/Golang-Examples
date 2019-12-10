package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/olivere/elastic/v7"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
)

func main() {
	flag.Parse()
	fmt.Println(flag.Args())
	client, err := elastic.NewClient(elastic.SetURL(flag.Arg(0), flag.Arg(1)))
	if err != nil {
		panic(err)
	}

	c, _ := client.Count("test_elastic15").Query(elastic.NewBoolQuery().Must(elastic.NewRangeQuery("Time").Gte(1563363900).Lte(1563364200))).Do(context.Background())
	fmt.Println(c)
	//qNum := count / 100000

	//threadNum:=10
	wg := sync.WaitGroup{}
	//f,_:=os.Create("r.json")
	//
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//	go func() {
	//		wg.Add(1)
	//		r,err:=client.Search("test_elastic15").Size(10000).From(i*100000).Query(
	//			elastic.NewBoolQuery().Must(elastic.NewRangeQuery("Time").Gte(1563363900).Lte(1563364200))).Do(context.Background())
	//		if err!=nil {
	//			fmt.Println(err)
	//		}else {
	//			for _,hit:=range r.Hits.Hits {
	//				f.Write(hit.Source)
	//			}
	//			fmt.Println(r.Error)
	//			fmt.Println(len(r.Hits.Hits),r.TotalHits())
	//		}
	//		wg.Done()
	//	}()
	//}

	//wg.Wait()
	strChan := make(chan string, 1000)

	//r, err := client.Scroll("test_elastic15").KeepAlive("10m").Size(100000).
	//	Query(elastic.NewBoolQuery().Must(elastic.NewRangeQuery("Time").Gte(1563363900).Lte(1563364200))).Do(context.Background())
	//if err != nil {
	//	fmt.Println(err)
	//}

	counter := int32(0)
	f, _ := os.Create("r.json")

	threadNum := 12
	wg.Add(threadNum)
	for i := 0; i < threadNum; i++ {
		go func(id int) {
			Slice := elastic.NewSliceQuery().Id(id).Max(threadNum)
			scroll := client.Scroll("test_elastic15").Query(
				elastic.NewBoolQuery().Must(elastic.NewRangeQuery("Time").Gte(1563363900).Lte(1563364200))).Slice(Slice).Size(100000)

			for {
				r, err := scroll.Do(context.Background())
				if err != nil {
					strChan <- "error : " + err.Error()
					if err.Error() == "EOF" {
						wg.Done()
						scroll.Clear(context.Background())
						break
					}
				} else {
					atomic.AddInt32(&counter, 1)
					strChan <- fmt.Sprintf("%v %v %v", atomic.LoadInt32(&counter), strconv.Itoa(len(r.Hits.Hits)), id)

					for _, hit := range r.Hits.Hits {
						_, err := f.Write(append(hit.Source, 10))
						if err != nil {
							strChan <- fmt.Sprintf("%v %v", id, err)
						}
					}
				}
			}
		}(i)
	}

	go func() {
		wg.Wait()
		os.Exit(0)
	}()

	for {
		fmt.Println(<-strChan)
	}
}
