package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Ackar/intro-to-go-workshop/internal/tunnel"
)

func main() {
	serverURL := "wss://workshop.sycl.dev/ws"
	port := 4242

	proxy, err := tunnel.NewWebsocketClientHTTPProxy(serverURL, port)
	if err != nil {
		log.Fatal(err)
	}
	go proxy.Run()

	http.HandleFunc("/info", InfoHandler)
	http.HandleFunc("/level1", Level1Handler)
	http.HandleFunc("/level2", NewLevel2().Handler)
	http.HandleFunc("/level3", Level3Handler)

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
