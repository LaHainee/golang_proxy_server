package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"

	"proxy_server/config"
	"proxy_server/internal/app/handlers"
	"proxy_server/internal/pkg/constants"
	"proxy_server/internal/pkg/utils"
)

func main() {
	// creating and parsing app config
	conf := config.NewConfig()
	if _, err := toml.DecodeFile(constants.ConfigPath, &conf); err != nil {
		logrus.Fatalf("Could not decode config file: %s", err)
	}

	// create logger object
	logger := utils.NewLogger(conf)

	// create the handlers
	proxyHandler := handlers.NewProxy(logger)

	// create a new serve mux and register the handlers
	serveMux := http.NewServeMux()
	serveMux.Handle("/", proxyHandler)

	// create a new server
	server := http.Server{
		Addr:         conf.ServerConfig.Port,
		Handler:      serveMux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// start the server
	go func() {
		logger.Printf("Starting server at %s", conf.ServerConfig.Port)

		err := server.ListenAndServe()
		if err != nil {
			logger.Errorf("Error starting server: %s", err)
			os.Exit(1)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-done
	logger.Println("Received terminate, graceful shutdown")

	// Graceful shutdown
	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(tc); err != nil {
		logger.Fatalf("Sever shutdown failed: %s", err)
	}
	logger.Println("Server exited properly")
}
