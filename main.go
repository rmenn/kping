package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var (
	version string
	healthy bool
	ready   bool
	delay   int
)

func pinghandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(version))
	fmt.Println("ponged a ping")
}

func healthyhandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		status := http.StatusOK
		if healthy == false {
			status = http.StatusTeapot
		}
		w.WriteHeader(status)
		w.Write([]byte("health"))
	case "POST":
		healthy = false
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(" health status false"))
	}
}

func readyhandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		status := http.StatusOK
		if ready == false {
			status = http.StatusTeapot
		}
		w.WriteHeader(status)
		w.Write([]byte("ready"))
	case "POST":
		ready = false
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(" ready status false"))
	}
}

func startuphandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		status := http.StatusOK
		time.Sleep(time.Duration(delay) * time.Second)
		w.WriteHeader(status)
		w.Write([]byte("started"))
	}
}

func main() {
	healthy = true
	ready = true
	delay = 1
	edelay, ok := os.LookupEnv("KPING_DELAY")
	if ok {
		delay, _ = strconv.Atoi(edelay)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", pinghandler)
	mux.HandleFunc("/healthz", healthyhandler)
	mux.HandleFunc("/ready", readyhandler)
	mux.HandleFunc("/start", startuphandler)
	srv := &http.Server{
		Addr:    ":80",
		Handler: mux,
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
		<-sigint
		healthy = false
		ready = false
		log.Println("shutting down")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("error shutting down : %v", err)
		}
		close(idleConnsClosed)
	}()

	log.Println("Starting HTTP server")
	log.Println("ping listening on 0.0.0.0, port 80")
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("error starting server : %v", err)
	}
	<-idleConnsClosed
}
