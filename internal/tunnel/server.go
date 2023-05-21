package tunnel

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/olahol/melody"
)

type WebsocketHTTPProxyServer struct {
	melody            *melody.Melody
	respChans         map[string]chan *HTTPResponse
	sessions          map[string]*melody.Session
	connectHandler    func(string)
	disconnectHandler func(string)
}

func NewWebsocketHTTPProxyServer() *WebsocketHTTPProxyServer {
	m := melody.New()
	m.Config.MaxMessageSize = 16000

	p := &WebsocketHTTPProxyServer{
		melody:            m,
		respChans:         make(map[string]chan *HTTPResponse),
		sessions:          make(map[string]*melody.Session),
		connectHandler:    func(s string) {},
		disconnectHandler: func(s string) {},
	}

	m.HandleMessage(p.HandleMessage)
	m.HandleConnect(func(s *melody.Session) {
		id := uuid.New().String()
		s.Set("id", id)
		p.sessions[id] = s
		go p.connectHandler(id)
	})

	m.HandleDisconnect(func(s *melody.Session) {
		id, _ := s.Keys["id"].(string)
		delete(p.sessions, id)
		go p.disconnectHandler(id)
	})

	return p
}

func (p *WebsocketHTTPProxyServer) SetConnectHandler(handler func(string)) {
	p.connectHandler = handler
}

func (p *WebsocketHTTPProxyServer) SetDisconnectHandler(handler func(string)) {
	p.disconnectHandler = handler
}

func (p *WebsocketHTTPProxyServer) WSHandler(w http.ResponseWriter, r *http.Request) {
	p.melody.HandleRequest(w, r)
}

func (p *WebsocketHTTPProxyServer) ExecuteRequest(sessionID string, req *http.Request) (*http.Response, error) {
	sess, ok := p.sessions[sessionID]
	if !ok {
		return nil, fmt.Errorf("unknown session")
	}

	proxyReq := FromHTTPRequest(req)
	reqJSON, _ := json.Marshal(proxyReq)
	sess.WriteBinary(reqJSON)

	respChan := make(chan *HTTPResponse)
	defer func() {
		close(respChan)
		delete(p.respChans, proxyReq.ID)
	}()
	p.respChans[proxyReq.ID] = respChan

	select {
	case resp := <-respChan:
		return resp.ToResponse(), nil
	case <-time.After(2 * time.Second):
		return nil, fmt.Errorf("request timeout")
	}
}

func (p *WebsocketHTTPProxyServer) HandleMessage(s *melody.Session, msg []byte) {
	var resp HTTPResponse
	json.Unmarshal(msg, &resp)

	respChan, ok := p.respChans[resp.ID]
	if !ok {
		return
	}
	respChan <- &resp
}
