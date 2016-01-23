// Simple live charts for memory consumption and GC pauses.
//
// To use debugcharts, link this package into your program:
//	import _ "github.com/mkevac/debugcharts"
//
// If your application is not already running an http server, you
// need to start one.  Add "net/http" and "log" to your imports and
// the following code to your main function:
//
// 	go func() {
// 		log.Println(http.ListenAndServe("localhost:6060", nil))
// 	}()
//
// Then go look at charts:
//
//	http://localhost:6060/debug/charts
//
package debugcharts

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mkevac/debugcharts/bindata"
	"github.com/shirou/gopsutil/process"
)

type update struct {
	Ts             int64
	BytesAllocated uint64
	GcPause        uint64
	CpuUser        float64
	CpuSys         float64
}

type consumer struct {
	id uint
	c  chan update
}

type server struct {
	consumers      []consumer
	consumersMutex sync.RWMutex
}

type SimplePair struct {
	Ts    uint64
	Value uint64
}

type CPUPair struct {
	Ts   uint64
	User float64
	Sys  float64
}

type DataStorage struct {
	BytesAllocated []SimplePair
	GcPauses       []SimplePair
	CpuUsage       []CPUPair
}

const (
	maxCount int = 86400
)

var (
	data           DataStorage
	lastPause      uint32
	mutex          sync.RWMutex
	lastConsumerId uint
	s              server
	upgrader       = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	prevSysTime  float64
	prevUserTime float64
	myProcess    *process.Process
)

func (s *server) gatherData() {
	timer := time.Tick(time.Second)

	for {
		select {
		case now := <-timer:
			nowUnix := now.Unix()

			var ms runtime.MemStats
			runtime.ReadMemStats(&ms)

			u := update{
				Ts: nowUnix * 1000,
			}

			cpuTimes, _ := myProcess.CPUTimes()

			if prevUserTime != 0 {
				u.CpuUser = cpuTimes.User - prevUserTime
				u.CpuSys = cpuTimes.System - prevSysTime
				data.CpuUsage = append(data.CpuUsage, CPUPair{uint64(nowUnix) * 1000, u.CpuUser, u.CpuSys})
			}

			prevUserTime = cpuTimes.User
			prevSysTime = cpuTimes.System

			mutex.Lock()

			bytesAllocated := ms.Alloc
			u.BytesAllocated = bytesAllocated
			data.BytesAllocated = append(data.BytesAllocated, SimplePair{uint64(nowUnix) * 1000, bytesAllocated})
			if lastPause == 0 || lastPause != ms.NumGC {
				gcPause := ms.PauseNs[(ms.NumGC+255)%256]
				u.GcPause = gcPause
				data.GcPauses = append(data.GcPauses, SimplePair{uint64(nowUnix) * 1000, gcPause})
				lastPause = ms.NumGC
			}

			if len(data.BytesAllocated) > maxCount {
				data.BytesAllocated = data.BytesAllocated[len(data.BytesAllocated)-maxCount:]
			}

			if len(data.GcPauses) > maxCount {
				data.GcPauses = data.GcPauses[len(data.GcPauses)-maxCount:]
			}

			mutex.Unlock()

			s.sendToConsumers(u)
		}
	}
}

func init() {
	http.HandleFunc("/debug/charts/data-feed", s.dataFeedHandler)
	http.HandleFunc("/debug/charts/data", dataHandler)
	http.HandleFunc("/debug/charts/", handleAsset("static/index.html"))
	http.HandleFunc("/debug/charts/main.js", handleAsset("static/main.js"))
	http.HandleFunc("/debug/charts/jquery-2.1.4.min.js", handleAsset("static/jquery-2.1.4.min.js"))
	http.HandleFunc("/debug/charts/moment.min.js", handleAsset("static/moment.min.js"))

	myProcess, _ = process.NewProcess(int32(os.Getpid()))

	// preallocate arrays in data, helps save on reallocations caused by append()
	// when maxCount is large
	data.BytesAllocated = make([]SimplePair, 0, maxCount)
	data.GcPauses = make([]SimplePair, 0, maxCount)
	data.CpuUsage = make([]CPUPair, 0, maxCount)

	go s.gatherData()
}

func (s *server) sendToConsumers(u update) {
	s.consumersMutex.RLock()
	defer s.consumersMutex.RUnlock()

	for _, c := range s.consumers {
		c.c <- u
	}
}

func (s *server) removeConsumer(id uint) {
	s.consumersMutex.Lock()
	defer s.consumersMutex.Unlock()

	var consumerId uint
	var consumerFound bool

	for i, c := range s.consumers {
		if c.id == id {
			consumerFound = true
			consumerId = uint(i)
			break
		}
	}

	if consumerFound {
		s.consumers = append(s.consumers[:consumerId], s.consumers[consumerId+1:]...)
	}
}

func (s *server) addConsumer() consumer {
	s.consumersMutex.Lock()
	defer s.consumersMutex.Unlock()

	lastConsumerId += 1

	c := consumer{
		id: lastConsumerId,
		c:  make(chan update),
	}

	s.consumers = append(s.consumers, c)

	return c
}

func (s *server) dataFeedHandler(w http.ResponseWriter, r *http.Request) {
	var (
		lastPing time.Time
		lastPong time.Time
	)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	conn.SetPongHandler(func(s string) error {
		lastPong = time.Now()
		return nil
	})

	// read and discard all messages
	go func(c *websocket.Conn) {
		for {
			if _, _, err := c.NextReader(); err != nil {
				c.Close()
				break
			}
		}
	}(conn)

	c := s.addConsumer()

	defer func() {
		s.removeConsumer(c.id)
		conn.Close()
	}()

	var i uint

	for u := range c.c {
		websocket.WriteJSON(conn, u)
		i += 1

		if i%10 == 0 {
			if diff := lastPing.Sub(lastPong); diff > time.Second*60 {
				return
			}
			now := time.Now()
			if err := conn.WriteControl(websocket.PingMessage, nil, now.Add(time.Second)); err != nil {
				return
			}
			lastPing = now
		}
	}
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	mutex.RLock()
	defer mutex.RUnlock()

	if e := r.ParseForm(); e != nil {
		log.Print("error parsing form")
		return
	}

	callback := r.FormValue("callback")

	fmt.Fprintf(w, "%v(", callback)

	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	encoder.Encode(data)

	fmt.Fprint(w, ")")
}

func handleAsset(path string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := bindata.Asset(path)
		if err != nil {
			log.Print(err)
			return
		}

		n, err := w.Write(data)
		if err != nil {
			log.Print(err)
			return
		}

		if n != len(data) {
			log.Print("wrote less than supposed to")
			return
		}
	}
}
