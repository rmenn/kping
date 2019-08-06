package main

import (
	"fmt"
	"net/http"
)

var (
	version string
	healthy bool
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
			status = http.StatusInternalServerError
		}
		w.WriteHeader(status)
		w.Write([]byte("health"))
	case "POST":
		healthy = false
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(" health status false"))
	}
}
func main() {
	healthy = true
	http.HandleFunc("/ping", pinghandler)
	http.HandleFunc("/healthz", healthyhandler)
	fmt.Println("ping listening on 0.0.0.0, port 80")
	err := http.ListenAndServe(":80", nil)

	if err != nil {
		fmt.Println("Error starting ping server: ", err)
	}
}
