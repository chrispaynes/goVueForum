package router

import (
	"context"
	"goVueForum/api/pkg/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

// New creates a new router with routes and handlerFuncs
func New() *mux.Router {
	return handle(mux.NewRouter())
}

func handle(r *mux.Router) *mux.Router {
	r.HandleFunc("/healthz", handlers.GetHealth).Methods("GET")
	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("OPTIONS", "POST")
	return r
}

// GracefulShutdown shuts the server down gracefully upon syscall notification
func GracefulShutdown(addr string) {
	srv := &http.Server{
		Addr: addr,
	}

	signalChan := make(chan os.Signal, 1)

	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	<-signalChan

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	srv.Shutdown(ctx)

	log.Println("Shutdown signal received, exiting...")
	os.Exit(0)
}
