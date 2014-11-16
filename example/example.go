package main

import (
	"log"
	"net/http"

	_ "expvar"

	_ "github.com/mkevac/debugcharts"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}
