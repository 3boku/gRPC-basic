package main

import (
	"flag"
	"gRPC-basic/config"
)

var configFlag = flag.String("config", "./config.toml", "config path")

func main() {
	flag.Parse()
	config.NewConfig(*configFlag)
}
