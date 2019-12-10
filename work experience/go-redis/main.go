package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

func main() {
	//redisdb := redis.NewClient(&redis.Options{
	//	Addr:         "localhost:6379",
	//	DialTimeout:  10 * time.Second,
	//	ReadTimeout:  30 * time.Second,
	//	WriteTimeout: 30 * time.Second,
	//	PoolSize:     10,
	//	PoolTimeout:  30 * time.Second,
	//	//Password:"12345678",
	//	DB:5,
	//})
	//
	//mymap:=make(map[string]interface{})
	//
	//mymap["username"]="amir"
	//mymap["mac"]="1,1,1"
	//
	//redisdb.HMSet("REBIN||",mymap)
	//
	//l, _ := redisdb.HMGet("REBIN||", "username", "mac").Result()
	//
	//for k,v:=range l{
	//	fmt.Println(k,v)
	//}
	//
	//fmt.Println(redisdb.HMGet("REBIN||", "username", "mac").Result())
	//redis.
	var p Producer
	p.connectToRedis()
	for i := 0; i < 10; i++ {
		m := make(map[string]interface{})
		m["1"] = "1"
		m["2"] = "2"
		fmt.Println(p.writeIPMap("test"+strconv.Itoa(i), m))
	}
}

type Producer struct {
	ip                string
	port              string
	redisPass         string
	dbNum             int
	threadCount       uint
	redisConnPoolSize int
	redisClient       *redis.Client
}

func (producer *Producer) connectToRedis() {
	producer.redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		//Password:     "",
		//DB:           config.Config.Producer.DBNum,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     30,
		PoolTimeout:  30 * time.Second,
	})
}

func (producer *Producer) writeIPMap(key string, Info map[string]interface{}) error {

	_, err := producer.redisClient.HMSet("REBIN||"+key, Info).Result()

	if err != nil {
		return fmt.Errorf("redis Set Error-> %s", err)
	}

	return nil
}
