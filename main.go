package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Ackar/intro-to-go-workshop/internal/tunnel"
)

var port = 4242

func main() {
	initProxy() // Proxy to workshop server, do not modify

	http.HandleFunc("/info", InfoHandler)
	http.HandleFunc("/level1", Level1Handler)
	http.HandleFunc("/level2", NewLevel2().Handler)
	http.HandleFunc("/level3", Level3Handler)

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func initProxy() {
	serverURL := "wss://workshop.sycl.dev/ws"

	proxy, err := tunnel.NewWebsocketClientHTTPProxy(serverURL, port)
	if err != nil {
		log.Fatal(err)
	}
	go proxy.Run()
}
