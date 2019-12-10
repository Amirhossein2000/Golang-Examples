package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/olivere/elastic/v7"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

type RebinOutput struct {
	Username      string
	Mac           string
	Time          int64
	ProtoNum      string
	InvalidIp     string
	InvalidPort   string
	ValidIp       string
	ValidPort     string
	DstIp         string
	DstPort       string
	Loc1          string
	Loc2          string
	Apn           string
	Imsi          string
	Imei          string
	OperatorName  string
	ServiceNum    int32
	CollectorName string
}

const mapping = `
{
	"settings" : {
		"number_of_shards" : 4,
		"number_of_replicas" : 0,
		"index" : {
			"codec": "best_compression",
			"max_result_window" : 120270946
		}
	},
	"mappings":{
			"properties":{
				"Username":{
					"type":"text",
					"index":false,
					"norms": false
				},
				"Mac":{
					"type":"text",
					"index":false,
					"norms": false
				},
				"Time":{
					"type":"date"
				},
				"ProtoNum":{
					"type":"text",
					"index":false,
					"norms": false
				},
				"InvalidIp":{
					"type":"text",
					"index":false,
					"norms": false
				},
				"InvalidPort":{
					"type":"text",
					"index":false,
					"norms": false
				},
				"ValidIp":{
					"type":"text",
					"norms": false
				},
				"ValidPort":{
					"type":"text",
					"index":false,
					"norms": false
				},
				"DstIp":{
					"type":"text",
					"index":false,
					"norms": false
				},
				"DstPort":{
					"type":"text",
					"index":false,
					"norms": false
				},
				"Loc1":{
					"type":"text",
					"index":false,
					"norms": false
				},
				"Loc2":{
					"type":"text",
					"index":false,
					"norms": false
				},
				"Apn":{
					"type":"text",
					"index":false,
					"norms": false
				},
				"Imei":{
					"type":"text",
					"index":false,
					"norms": false
				},
				"OperatorName":{
					"type":"text",
					"index":false,
					"norms": false
				},
				"ServiceNum":{
					"type":"text",
					"norms": false
				},
				"CollectorName":{
					"type":"text",
					"norms": false
				}
			}
		}
}
`

var logch chan string

func main() {
	logch = make(chan string, 1000)
	//pvr := make(map[int]string)
	//pvr[0] = "tehran"
	//pvr[1] = "AzarbayjanEst"
	//pvr[2] = "Semnan"
	//pvr[3] = "Elam"
	//pvr[5] = "ESFAHAN"
	//pvr[6] = "Elam"
	//pvr[7] = "Shiraz"

	flag.Parse()
	worker, _ := strconv.Atoi(flag.Arg(0))
	wc, _ := strconv.ParseInt(flag.Arg(1), 10, 64)
	batchSize := 1524
	indexName := "ipdr2"

	fmt.Println(flag.Args())

	client, err := elastic.NewClient(elastic.SetURL(flag.Arg(2), flag.Arg(3)))

	//client, err := elastic.NewClientFromConfig(&config.Config{
	//	//URL: "http://172.16.4.4:9200",
	//	URL: flag.Arg(2),
	//})

	if err != nil {
		panic(err)
	}

	_, _ = client.DeleteIndex(indexName).Do(context.Background())

	time.Sleep(time.Second)

	_, err = client.CreateIndex(indexName).BodyString(mapping).Do(context.Background())
	if err != nil {
		panic(err)
	}

	wg := sync.WaitGroup{}
	rand.Seed(time.Now().UnixNano())

	time.Sleep(time.Second * 5)
	allTime := time.Now()
	myTime := time.Now().Unix()

	for j := 0; j < worker; j++ {
		go func(n int) {
			wg.Add(1)
			defer wg.Done()

			bulkService := client.Bulk()

			for writeCount := int64(0); writeCount < wc; writeCount++ {
				for i := 0; i < batchSize; i++ {
					bulkService = bulkService.Add(elastic.NewBulkIndexRequest().Index(indexName).Doc(RebinOutput{
						Username:      "89e810932jksdh3e982",
						CollectorName: "Shiraz",
						Apn:           "apn",
						DstIp:         "1.1.1.1",
						DstPort:       "1080",
						Imei:          "aa:aa:aa:aa:aa:aa:aa:aa",
						Imsi:          "aa:aa:aa:aa:aa:aa:aa:aa",
						InvalidIp:     "8.8.8.8",
						InvalidPort:   "8080",
						Loc1:          "loc1",
						Loc2:          "loc2",
						Mac:           "99:11:11:22:33:33",
						OperatorName:  "tci",
						ProtoNum:      "9",
						ServiceNum:    10,
						Time:          myTime,
						ValidIp:       fmt.Sprintf("%v.%v.%v.%v", n, n, n, n),
						ValidPort:     "7070",
					}))
				}

				res, err := bulkService.Do(context.Background())
				if err != nil {
					log(fmt.Sprint(n, err.Error()))
					return
				}
				if res.Errors || len(res.Succeeded()) < batchSize {
					log(fmt.Sprint("there is an error", res.Errors, len(res.Succeeded()), res.Failed()[0].Status, res.Failed()[0].Error))
				}

				log(fmt.Sprint("end", n, writeCount))
			}

			//wg.Done()
		}(j)
	}

	go func() {
		time.Sleep(10 * time.Second)
		wg.Wait()
		fmt.Println("allTime:", time.Since(allTime))
		os.Exit(0)
	}()

	fmt.Println("logging")

}

var IndexName string
var MT sync.Mutex

func GetIndex() string {
	t := time.Now().Format("20060102")
	MT.Lock()
	if t != IndexName {
		IndexName = t
	}
	MT.Unlock()
	return IndexName
}

func log(s string) {
	logch <- s
}
