package main

import (
	"github.com/2336260845/hippo_recall/config"
	"github.com/2336260845/hippo_recall/es"
	"github.com/2336260845/hippo_recall/server"
)

func init() {
	config.InitConfig("")
	cf := config.GetConfig()
	es.InitEsClient(cf)
}

func main() {
	cf := config.GetConfig()
	server.ThriftInit(cf)
}
