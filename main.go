package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tianye0718/go-cache/gocache"
)

var db = map[string]string{
	"Tom": "630",
	"Bob": "621",
	"Sam": "578",
}

func main() {
	gocache.NewGroup("scores", 2<<10, gocache.GetterFunc(func(key string) ([]byte, error) {
		log.Println("[SlowDB] search key", key)
		if v, ok := db[key]; ok {
			return []byte(v), nil
		}
		return nil, fmt.Errorf("%s not exist", key)
	}))

	addr := "localhost:9999"
	peers := gocache.NewHTTPPool(addr)
	log.Println("gocache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}
