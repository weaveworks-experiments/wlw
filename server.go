package main

import (
	"log"
	"net/http"

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
	requestProcessingTimes = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name: "who_lives_where_processing_time",
			Help: "Processing time latencies for the who_lives_where service.",
			// 0.001 seconds = 1 millisecond, up to 10 seconds
			Buckets: prometheus.LinearBuckets(0.001, 0.001, 10000),
		},
		[]string{"who"},
	)
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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
	})
	http.HandleFunc("/person", func(w http.ResponseWriter, r *http.Request) {
		startTime = time.Now().UnixNano()
		person := r.URL.Query()["who"][0]
		responseCode := "200"
		if person == "anita" {
			fmt.Fprintf(w, "Anita lives in Mexico")
		} else if person == "tamao" {
			fmt.Fprintf(w, "Tamao lives in California")
		} else if person == "ilya" {
			time.Sleep(1 * time.Second)
			fmt.Fprintf("Ilya lives in London")
		} else {
			fmt.Fprintf(w, "Person not found, error!")
			responseCode = "404"
			w.WriteHeader(http.StatusNotFound)
		}
		timeNanoseconds = time.Now().UnixNano() - startTime
		requestTotals.WithLabelValues(responseCode, request.Method).Inc()
		requestProcessingTimes.WithLabelValues(person).Observe(
			//                          nano   micro  milli
			float64(timeNanoseconds) / (1000 * 1000 * 1000),
		)
	})

	prometheus.MustRegister(requestLatencies)
	prometheus.MustRegister(requestProcessingTimes)

	go func() {
		http.ListenAndServe(":8000", nil)
	}()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
