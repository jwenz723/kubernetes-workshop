package main

import (
	"flag"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io"
	"log"
	"net/http"
	"strings"
)

var metricRequestsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name:        "http_requests_total",
		Help:        "The total number of http requests received",
		ConstLabels: prometheus.Labels{"component": "backend"},
	},
	[]string{"method", "endpoint"},
)

const HELLO_RESPONSE = "Hello, from the underworld!\n"

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(metricRequestsTotal)
}

func main() {
	addr := flag.String("addr", ":8081", "address to run the backend server on")
	flag.Parse()

	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/contributors", handleContributors(metricRequestsTotal.WithLabelValues("GET", "/contributors")))
	http.Handle("/hello", handleHello(metricRequestsTotal.WithLabelValues("GET", "/hello")))
	fmt.Printf("starting backend http on %s\n", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func handleHello(requestsCounter prometheus.Counter) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		requestsCounter.Inc()
		fmt.Printf("handled hello request\n")
		io.WriteString(w, HELLO_RESPONSE)
	}
}

func handleContributors(requestsCounter prometheus.Counter) http.HandlerFunc {
	contributors := []string{"Jeff Wenzbauer"}
	return func(w http.ResponseWriter, req *http.Request) {
		requestsCounter.Inc()
		fmt.Printf("handled contributors request\n")
		io.WriteString(w, strings.Join(contributors, "\n"))
	}
}
