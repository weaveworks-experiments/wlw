package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestTotals = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "who_lives_where_request_count",
			Help: "Count of requests for who_lives_where service.",
		},
		[]string{"code", "method"},
	)
	requestProcessingTimes = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "who_lives_where_processing_time",
			Help: "Processing time latencies for the who_lives_where service.",
			// 0.1 seconds = 100 millisecond, up to 10 seconds
			Buckets: prometheus.LinearBuckets(0.1, 0.1, 100),
		},
		[]string{"who"},
	)
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now().UnixNano()
		fmt.Fprintf(w, `
		<html><head><title>Who lives where?</title></head>
		<body>
		<p>Hi there!</p>
		<p>Find out who lives where:</p>
		<ul>
			<li><a href="/person?who=anita">Anita</a></li>
			<li><a href="/person?who=tamao">Tamao</a></li>
			<li><a href="/person?who=ilya">Ilya</a></li>
		</ul>
		</body>
		</html>
		`)
		timeNanoseconds := time.Now().UnixNano() - startTime
		requestTotals.WithLabelValues("200", r.Method).Inc()
		requestProcessingTimes.WithLabelValues("home").Observe(
			//                          micro  milli  seconds
			float64(timeNanoseconds) / (1000 * 1000 * 1000),
		)
	})
	http.HandleFunc("/person", func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now().UnixNano()
		person := r.URL.Query()["who"][0]
		responseCode := "200"
		if person == "anita" {
			fmt.Fprintf(w, "Anita lives in Mexico")
		} else if person == "tamao" {
			fmt.Fprintf(w, "Tamao lives in California")
		} else if person == "ilya" {
			fmt.Fprintf(w, "Ilya lives in London")
		} else {
			fmt.Fprintf(w, "Person not found, error!")
			responseCode = "404"
			w.WriteHeader(http.StatusNotFound)
		}
		timeNanoseconds := time.Now().UnixNano() - startTime
		requestTotals.WithLabelValues(responseCode, r.Method).Inc()
		requestProcessingTimes.WithLabelValues(person).Observe(
			//                          micro  milli  seconds
			float64(timeNanoseconds) / (1000 * 1000 * 1000),
		)
	})

	prometheus.MustRegister(requestTotals)
	prometheus.MustRegister(requestProcessingTimes)

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
