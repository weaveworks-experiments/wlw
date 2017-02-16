package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	req := func() {
		_, err := http.Get("http://localhost:30003/")
		if err != nil {
			log.Fatal(err)
		}
		log.Print(".")
	}
	// make some requests with a spike
	// aim to simulate (ish) 1, 2, 3, 13, 23, 33, 34, 35, 36
	for i := 0; i < 6; i++ {
		req()
		time.Sleep(1 * time.Second)
	}
	for i := 0; i < 30; i++ {
		req()
		time.Sleep(100 * time.Millisecond)
	}
	for i := 0; i < 6; i++ {
		req()
		time.Sleep(1 * time.Second)
	}
}
