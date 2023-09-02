package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/zuri03/TimeCapsule/capsule"
	"go.uber.org/zap"
)

func main() {
	fmt.Println("Starting Server...")
	requestWaitGroup := new(sync.WaitGroup)
	baseLogger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("Unable to intialize zap logger: %s\n", err.Error())
	}
	logger := baseLogger.Sugar()
	middleware := capsule.Pipeline(requestWaitGroup, logger)
	service := capsule.Service(logger)
	controller := capsule.Controller(service, logger)

	address := ":9000"
	server := &http.Server{
		Addr:        address,
		Handler:     middleware.ValidateRequest(middleware.ParseCapsuleFromRequest(controller)),
		IdleTimeout: 10 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Println("Server now listening on port 9000")
	signaler := make(chan os.Signal)
	signal.Notify(signaler, os.Interrupt)
	signal.Notify(signaler, os.Kill)

	<-signaler

	fmt.Println("Shutdown signal received, waiting for existing requests to finish...")

	requestWaitGroup.Wait()

	fmt.Println("Exiting...")
}
