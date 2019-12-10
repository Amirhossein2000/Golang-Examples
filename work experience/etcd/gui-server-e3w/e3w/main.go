package main

import (
	"flag"
	"fmt"
	"github.com/coreos/etcd/version"
	"github.com/gin-gonic/gin"
	"github.com/soyking/e3w/conf"
	"github.com/soyking/e3w/e3ch"
	"github.com/soyking/e3w/routers"
	"os"
)

const (
	PROGRAM_NAME    = "e3w"
	PROGRAM_VERSION = "0.0.2"
)

var configFilepath string

func init() {
	flag.StringVar(&configFilepath, "conf", "conf/config.default.ini", "config file path")
	rev := flag.Bool("rev", false, "print rev")
	flag.Parse()

	if *rev {
		fmt.Printf("[%s v%s]\n[etcd %s]\n",
			PROGRAM_NAME, PROGRAM_VERSION,
			version.Version,
		)
		os.Exit(0)
	}
}

func main() {
	fmt.Println("asgdjkasdkjkads")
	config, err := conf.Init(configFilepath)
	fmt.Println("asdasd")
	if err != nil {
		panic(err)
	}

	client, err := e3ch.NewE3chClient(config)
	if err != nil {
		panic(err)
	}
	fmt.Println("asdasd222")

	router := gin.Default()
	router.UseRawPath = true
	routers.InitRouters(router, config, client)
	fmt.Println("config",config.Port)
	router.Run(":" + config.Port)
}
