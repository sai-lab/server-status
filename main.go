package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"./lib"
)

func handler(w http.ResponseWriter, r *http.Request) {

	var buf bytes.Buffer
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			d, err := lib.GetServerStat()
			if err != nil {
				log.Println("get status error:", err)
				d.ErrorInfo = err
			}
			j, _ := json.Marshal(d)
			buf.Write(j)
			fmt.Fprintf(w, buf.String())

		case <-interrupt:
			log.Println("interrupt")
			return
		}
	}

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
