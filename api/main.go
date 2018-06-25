package main

import (
	"goVueForum/api/pkg/postgres"
	"goVueForum/api/pkg/rabbitmq"
	"goVueForum/api/pkg/router"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/urfave/negroni"
)

var log = logrus.New()

func main() {
	log.Formatter = &logrus.JSONFormatter{}

	viper.SetConfigName("app")
	viper.AddConfigPath(".")
	viper.WatchConfig()

	if err := viper.ReadInConfig(); err != nil {
		log.Infof("failed to read config file: %s", err.Error())
	}

	r := router.New()
	log.Infoln("Starting goVueForum application...")

	rabbitURL := viper.GetString("dev.rabbitMQ_url")
	port := viper.GetString("dev.server_port")
	addr := viper.GetString("dev.server_address") + port

	if rabbitURL == "amqp://user:password@0.0.0.0:5672" {
		log.Fatalln("invalid rabbitMQ connection string")
	}

	defer router.GracefulShutdown(addr)

	n := negroni.Classic()
	n.UseHandler(r)

	go func() {
		log.Infof("goVueForumapplication is listening @ %s\n", addr)
		log.Fatalln(http.ListenAndServe(port, r))
	}()

	rabbitmq.Init(rabbitURL)

	c := postgres.NewConn(viper.GetString("dev.bridge_IP"))
	err := c.Open()

	if err != nil {
		log.Infof("failed to open Postgres database connection: %v", err)
	}

	c.GetCategories()

	defer c.Close()
}
