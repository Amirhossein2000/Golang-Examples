package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/etcdserver/api/v3rpc/rpctypes"
	"log"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	//defer cli.Close()

	ctx, _ := context.WithTimeout(context.Background(), time.Second)

	resp, err := cli.Put(ctx, "rebin", `[GENERAL]
operator_name="PP"
read_location_from_file=false

[SERVICE_IP]
default_service="ADSL"
DIALUP=[]
ADSL=[]
WIRELESS=[]
TD_LTE=[]
WIFI=[]
MOBILE_2G=[]
MOBILE_3G=[]
MOBILE_4G=[]
MOBILE_5G=[]
DEDICATED_BANDWIDTH=[]

[LOGGING]
#Leave destination empty for local syslog. Otherwise IP:Port
syslog_destination=""

[IPMAP_SOURCE]
source="IBS"
empty_username=false
empty_service_number=false
redis_host="127.0.0.1"
redis_port=6379
redis_db=5
redis_pass=""
redis_pool_size=256

[CSV_IPMAP]
csv_path=""

[CSV_APN]
csv_path=""
redis_host="127.0.0.1"
redis_port=6379
redis_db=6
redis_pass=""
redis_pool_size=256

[STATISTICS]
stat_agentx_master_net="unix"
stat_agentx_master_address="/var/agentx/master"

[OUTPUT_DESTINATION]
destination="SYSLOG"

[SYSLOG_OUTPUT]
#IP:Port to send output data
syslog_destination="127.0.0.1:515"

[KAFKA_OUTPUT]
broker="127.0.0.1"
group="rebin_log_producers"
topic="rebin_output"

[CONFIG_MANAGER]
host="127.0.0.1"
port=2379
password="123"
username="rebin-user"
rebin_key="rebin"

#[[KAFKA_SERVER]]
#this is the sample
#broker= "172.16.4.4"
#group= "dataflow_consumers"
#topic="rebin_dataflow"
#name= "ParsPooyesh"
#weight= 2
#workers=64

[[KAFKA_SERVER]]
broker= ""
group= ""
topic=""
name= ""
weight= 1
workers=32`)

	//	resp, err := cli.Put(ctx, "rebin-writer", `[GENERAL]
	//#Number of concurrent pipes. Ideally cpu cores/4
	//concurrency=2
	//
	//[LOGGING]
	//#Leave destination empty for local syslog. Otherwise IP:Port
	//syslog_destination=""
	//
	//[STATISTICS]
	//stat_agentx_master_net="unix"
	//stat_agentx_master_address="/var/agentx/master"
	//
	//[KAFKA]
	//broker="172.16.4.4"
	//group="rebin_output_consumers"
	//topic="rebin_output"
	//consumers_count=1
	//
	//[OUTPUT_DESTINATION]
	//destination="CASSANDRA,SYSLOG"
	//
	//[SYSLOG_OUTPUT]
	//#IP:Port to send output data
	//syslog_destination="127.0.0.1:515"
	//
	//[Cassandra_Output]
	//cassandra_destination=["127.0.0.1"]
	//key_space="tci_log"
	//table_name="ipdr"
	//
	//[MONGO_OUTPUT]
	//#IP:Port to send output data
	//#youruser:yourpassword@IP:PORT/yourdatabase for authentication
	//mongo_destination="127.0.0.1:27017"
	//database="rebin"
	//collection="ipdr"
	//mongo_batch_size=1024
	//
	//[CONFIG_MANAGER]
	//host="127.0.0.1"
	//port=2379
	//password="123"
	//username="rebin-user"
	//rebin_key="rebin"`)

	fmt.Println(resp)

	//cli.Put(ctx,"dir1/1","a")
	//cli.Put(ctx,"dir1/2","b")
	//cli.Put(ctx,"dir1/3","c")

	resp1, _ := cli.Get(ctx, "dir123", clientv3.WithPrefix())

	for _, kv := range resp1.Kvs {
		fmt.Println(string(kv.Value))
	}

	if err != nil {
		switch err {
		case context.Canceled:
			log.Fatalf("ctx is canceled by another routine: %v", err)
		case context.DeadlineExceeded:
			log.Fatalf("ctx is attached with a deadline is exceeded: %v", err)
		case rpctypes.ErrEmptyKey:
			log.Fatalf("client-side error: %v", err)
		default:
			log.Fatalf("bad cluster endpoints, which are not etcd servers: %v", err)
		}
	}
}
