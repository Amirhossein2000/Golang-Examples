package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

func main() {
	client, err := elastic.NewClient(elastic.SetURL("http://172.16.4.6:9200"))

	if err != nil {
		panic(err)
	}

	body := `
{
  "index_patterns": ["te*"],
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
					"index":false,
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
	body=""
	fmt.Println(body)
	//r,err:=elastic.NewIndicesPutTemplateService(client).Name("template").Create(true).BodyString(body).Do(context.Background())

	r, err := elastic.NewIndicesGetTemplateService(client).Name("template").Pretty(true).Do(context.Background())
	fmt.Println(err)

	fmt.Println(r["template"])

	client.CreateIndex(r["template"].)
}
