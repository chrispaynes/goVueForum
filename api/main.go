package main

import (
	"goVueForum/api/pkg/router"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

var log = logrus.New()

func main() {
	log.Formatter = &logrus.JSONFormatter{}

	r := router.New()
	log.Infoln("Starting goVueForum application...")

	addr := "0.0.0.0:3000"
	defer router.GracefulShutdown(addr)

	n := negroni.Classic()
	n.UseHandler(r)

	go func() {
		log.Infoln("goVueForumapplication is listening on Port 3000")
		log.Fatalln(http.ListenAndServe(":3000", r))
	}()

}
