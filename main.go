package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

const (
	// NoExpiration is for use to init Items without expiration, if it is -1, the item never expired.
	NoExpiration time.Duration = -1
	// DefaultExpiration is for use to init Items with expiration, if using 0, it used the configed expiration time.
	DefaultExpiration time.Duration = 0
)

func main() {
	store := Store{DefaultExpiration: 100, Items: make(map[string]Item), Lock: sync.RWMutex{}}
	http.HandleFunc("/", store.requestHandler)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
