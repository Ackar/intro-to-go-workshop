package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/Ackar/intro-to-go-workshop/internal/tunnel"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/rs/cors"

	_ "embed"
)

//go:embed schema.graphql
var schema string

func main() {
	proxy := tunnel.NewWebsocketHTTPProxyServer()
	clientManager := NewClientManager(proxy)
	levelHandler := NewLevelHandler(clientManager)

	frontendURL, err := url.Parse("http://localhost:3000/")
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", httputil.NewSingleHostReverseProxy(frontendURL))
	http.HandleFunc("/ws", proxy.WSHandler)

	schema := graphql.MustParseSchema(schema, levelHandler, graphql.UseFieldResolvers())
	http.Handle("/query", cors.Default().Handler(&relay.Handler{Schema: schema}))

	log.Println("Listening...")
	http.ListenAndServe(":8383", nil)
}
