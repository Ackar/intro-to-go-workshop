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

	// FIXME: Uncomment for level 2, step 3
	// randomColor := fmt.Sprintf("rgb(%d, %d, %d)", rand.Int31n(255), rand.Int31n(255), rand.Int31n(255))

	http.HandleFunc("/info", InfoHandler)
	http.HandleFunc("/level1", Level1Handler)
	http.HandleFunc("/level2", NewLevel2( /* FIXME (Level 2, Step 3): use the randomColor var (see above) */ ).Handler)
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
