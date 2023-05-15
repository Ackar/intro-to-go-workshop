package tunnel

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

type WebsocketClientHTTPProxy struct {
	localPort int
	serverURL string
}

func NewWebsocketClientHTTPProxy(serverURL string, localPort int) (*WebsocketClientHTTPProxy, error) {
	return &WebsocketClientHTTPProxy{
		serverURL: serverURL,
		localPort: localPort,
	}, nil
}

func (p *WebsocketClientHTTPProxy) Run() {
	for {
		conn, _, err := websocket.DefaultDialer.Dial(p.serverURL, nil)
		if err != nil {
			fmt.Printf("\033[31m❌ Couldn't connect to server (%v), will retry...\033[0m\n", err)
			time.Sleep(2 * time.Second)
			continue
		}
		fmt.Println("\033[32m🟢 Connected to server, you're ready to Go!\033[0m")

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}

			var req HTTPRequest
			err = json.Unmarshal(message, &req)
			if err != nil {
				return
			}
			fmt.Println("executing request", req.URL)
			u, _ := url.Parse(req.URL)
			u.Host = fmt.Sprintf("localhost:%d", p.localPort)
			u.Scheme = "http"
			r := &http.Request{
				Method: req.Method,
				Header: req.Header,
				URL:    u,
				Body:   io.NopCloser(bytes.NewReader(req.Body)),
			}

			resp, err := http.DefaultClient.Do(r)
			if err != nil {
				return
			}

			tunnelResp := FromHTTPResponse(req.ID, resp)

			conn.WriteJSON(tunnelResp)
		}
	}
}
