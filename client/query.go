package query

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	req := func() {
		resp, err := http.Get("http://localhost/")
		if err != nil {
			log.Fatal(err)
		}
	}
	// make some requests with a spike
	// aim to simulate 1, 2, 3, 13, 23, 33, 34, 35, 36
	for i = 0; i < 3; i++ {
		req()
		time.Sleep(1 * time.Second)
	}
	for i = 0; i < 30; i++ {
		req()
		time.Sleep(0.1 * time.Second)
	}
	for i = 0; i < 3; i++ {
		req()
		time.Sleep(1 * time.Second)
	}
}
