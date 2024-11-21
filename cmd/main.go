package main

import (
	"context"
	"demo/app/internal/auth"
	cfg "demo/app/internal/config"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	user := auth.User{}
	r := user.Route()

	mainRouter := mux.NewRouter()
	// merge all routes
	mainRouter.PathPrefix("/").Handler(user.Route())

	config := cfg.Config{}
	configData := config.Init()

	srv := &http.Server{
		Addr: "0.0.0.0:" + configData.Port,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	go func() {
		fmt.Println("Server is running on port:", configData.Port)
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatalf("server failed to start: %v", err)
		}

	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("shutting down")
	os.Exit(0)
}

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello, World!"))
}