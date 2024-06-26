package main

import (
	"flag"
	"gRPC-basic/cmd"
	"gRPC-basic/config"
)

var configFlag = flag.String("config", "./config.toml", "config path")

func main() {
	flag.Parse()
	cfg := config.NewConfig(*configFlag)
	cmd.NewApp(cfg)
}
