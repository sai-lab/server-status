package functions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/sai-lab/server-status/lib/status"
)

var loadCh = make(chan string, 1)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, <-loadCh)
}

func LoadMeasurement() {
	var buf bytes.Buffer
	log.SetFlags(0)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			d, err := status.GetServerStat()
			if err != nil {
				log.Println("get status error:", err)
				d.ErrorInfo = err
			}
			j, _ := json.Marshal(d)
			buf.Write(j)
			loadCh <- buf.String()
		}
	}
}
