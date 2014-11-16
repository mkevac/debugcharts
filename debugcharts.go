package debugcharts

import (
	"log"
	"net/http"
)

func init() {
	http.HandleFunc("/debug/charts/", handleAsset("static/index.html"))
	http.HandleFunc("/debug/charts/main.js", handleAsset("static/main.js"))
	/*
		http.HandleFunc("/debug/charts/proxy", handleProxy)
	*/
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

/*
func handleProxy(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://bumpd.mlan:21104/debug/vars")
	if err != nil {
		log.Println(err)
		// return something
	}

	headers := w.Header()
	headers.Set("Content-Type", "application/json")
	headers.Set("Access-Control-Allow-Origin", "*")

	io.Copy(w, resp.Body)
}
*/
