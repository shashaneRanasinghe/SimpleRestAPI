package main

import (
	"github.com/shashaneRanasinghe/simpleAPI/internal/config"
	"github.com/shashaneRanasinghe/simpleAPI/pkg/server"
	"github.com/tryfix/log"
)

func main() {

	config.LoadConfigs()
	closeChannel := server.Serve()
	<-closeChannel

	log.Info("Service Stopped")
}
