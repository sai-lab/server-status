package main

import (
	"net/http"
	"os"
	"os/signal"

	"github.com/sai-lab/server-status/lib/functions"
)

func main() {

	go functions.LoadMeasurement()

	http.HandleFunc("/", functions.Handler)
	http.ListenAndServe(":8080", nil)

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)

	for range channel {
		os.Exit(0)
	}
}
