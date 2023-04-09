package main

import (
	"200lab/common"
	"200lab/config"
	"200lab/internal/server"
	"200lab/sdk/httpserver"
	mMysql "200lab/sdk/mysql"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := mMysql.New(conf.Mysql.Dns)
	if err != nil {
		log.Fatal(err)
	}

	appContext := common.NewAppContext(conf, db)

	httpHandler := server.NewHttpHandler(appContext)

	server := httpserver.New(httpHandler, httpserver.WithAddress(conf.App.HttpHost, conf.App.HttpPort))
	server.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	server.Shutdown()
}
