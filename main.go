package main

import (
	"net/http"
	"os"
	"os/signal"
	"runtime"

	"github.com/sai-lab/server-status/lib/functions"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// cmd := functions.MakeDstatLog()

	http.HandleFunc("/", functions.Handler)
	http.ListenAndServe(":8080", nil)

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)

	for range channel {
		// cmd.Process.Kill()
		os.Exit(0)
	}
}
