package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var histogram = promauto.NewHistogram(prometheus.HistogramOpts{
	Name:    "random_numbers",
	Help:    "A histogram of normally distributed random numbers.",
	Buckets: prometheus.LinearBuckets(-3, .1, 61),
})

func main() {
	println("hello 8888")
	go func() {
		log.Fatal(http.ListenAndServe(":8181", promhttp.Handler()))
	}()
	mux := http.NewServeMux()
	mux.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("received %s", r.URL.Path)
		fmt.Fprintf(w, "Hello 2, %q", html.EscapeString(r.URL.Path))
		log.Printf("performed some action on %s", r.URL.Path)
	})
	log.Fatal(http.ListenAndServe(":8888", mux))
}
