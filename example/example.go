package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"

	_ "net/http/pprof"

	"github.com/gorilla/handlers"
	_ "github.com/mkevac/debugcharts"
)

func dummyCPUUsage() {
	var a uint64
	var t = time.Now()
	for {
		t = time.Now()
		a += uint64(t.Unix())
	}
}

func dummyAllocations() {
	var d []uint64

	for {
		for i := 0; i < 2*1024*1024; i++ {
			d = append(d, 42)
		}
		time.Sleep(time.Second * 10)
		fmt.Println(len(d))
		d = make([]uint64, 0)
		runtime.GC()
		time.Sleep(time.Second * 10)
	}
}

func main() {
	go dummyAllocations()
	go dummyCPUUsage()
	go func() {
		log.Fatal(http.ListenAndServe(":8080", handlers.CompressHandler(http.DefaultServeMux)))
	}()
	log.Printf("you can now open http://localhost:8080/debug/charts/ in your browser")
	select {}
}
