package main

import (
	"net/http"

	_ "github.com/lib/pq"

	"goVueForum/api/pkg/rabbitmq"
	"goVueForum/api/pkg/router"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/urfave/negroni"
)

var log = logrus.New()

func main() {
	log.Formatter = &logrus.JSONFormatter{}

	viper.SetConfigName("app")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Infof("failed to read config file: %s", err.Error())
	}

	r := router.New()
	log.Infoln("Starting goVueForum application...")

	rabbitURL := viper.GetString("dev.RABBITMQ_URL")
	port := viper.GetString("dev.SERVER_PORT")
	addr := viper.GetString("dev.SERVER_ADDRESS") + port

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
}
