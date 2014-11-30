package debugcharts

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime"
	"sync"
	"time"
)

type timestampedDatum struct {
	Count uint64 `json:"C"`
	Ts    int64  `json:"T"`
}

type exportedData struct {
	BytesAllocated []timestampedDatum
	GcPauses       []timestampedDatum
}

const (
	maxCount int = 86400
)

var (
	data      exportedData
	lastPause uint32
	mutex     sync.RWMutex
)

func gatherData() {
	timer := time.Tick(time.Second)

	for {
		select {
		case now := <-timer:
			nowUnix := now.Unix()
			var ms runtime.MemStats
			runtime.ReadMemStats(&ms)

			mutex.Lock()
			data.BytesAllocated = append(data.BytesAllocated, timestampedDatum{Count: ms.Alloc, Ts: nowUnix})
			if lastPause == 0 || lastPause != ms.NumGC {
				data.GcPauses = append(data.GcPauses, timestampedDatum{Count: ms.PauseNs[(ms.NumGC+255)%256], Ts: nowUnix})
				lastPause = ms.NumGC
			}

			if len(data.BytesAllocated) > maxCount {
				data.BytesAllocated = data.BytesAllocated[len(data.BytesAllocated)-maxCount:]
			}

			if len(data.GcPauses) > maxCount {
				data.GcPauses = data.GcPauses[len(data.GcPauses)-maxCount:]
			}

			mutex.Unlock()
		}
	}
}

func init() {
	/*
		http.HandleFunc("/debug/charts/data-feed", websocket.Handler(dataFeedHandler))
	*/
	http.HandleFunc("/debug/charts/data", dataHandler)
	http.HandleFunc("/debug/charts/", handleAsset("static/index.html"))
	http.HandleFunc("/debug/charts/main.js", handleAsset("static/main.js"))

	go gatherData()
}

/*
func (s server) dataFeedHandler(ws *websocket.Conn) {
	defer ws.Close()

	for {
		select {
		}
	}
}
*/

func dataHandler(w http.ResponseWriter, r *http.Request) {
	mutex.RLock()
	defer mutex.RUnlock()

	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func handleAsset(path string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := Asset(path)
		if err != nil {
			log.Fatal(err)
		}

		n, err := w.Write(data)
		if err != nil {
			log.Fatal(err)
		}

		if n != len(data) {
			log.Fatal("wrote less than supposed to")
		}
	}
}
