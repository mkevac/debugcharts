package main

import (
	"log"
	"net/http"
	"time"

	_ "net/http/pprof"

	"github.com/gorilla/handlers"
	_ "github.com/mkevac/debugcharts"
)

func dummyAllocations() {
	type t struct {
		a uint64
		b map[uint64][]byte
	}
	d := make([]t, 0, 0)

	for {
		select {
		case <-time.Tick(time.Second * 20):
			d = make([]t, 0, 0)
			log.Println(len(d))
		case <-time.Tick(time.Second * 8):
			d = make([]t, 500000, 500000)
			log.Println(len(d))
		}
	}
}

func main() {
	go dummyAllocations()
	log.Fatal(http.ListenAndServe(":8080", handlers.CompressHandler(http.DefaultServeMux)))
}
