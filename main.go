package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/goocarry/mstemplate/handlers"
)

func main() {
	log.Println("Hello microservices")

	logger := log.New(os.Stdout, "api", log.LstdFlags)
	helloHandler := handlers.NewHello(logger)

	sm := http.NewServeMux()
	sm.Handle("/", helloHandler)

	s := http.Server{
		Addr:         ":9990",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			logger.Fatalf("server error: %s", err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	logger.Println("Recieved terminate, graceful shutdown", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	s.Shutdown(tc)
}
