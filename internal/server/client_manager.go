package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/Ackar/intro-to-go-workshop/internal/tunnel"
)

type ClientInfo struct {
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
}

type ClientManager struct {
	proxyServer *tunnel.WebsocketHTTPProxyServer
	clientsInfo map[string]*ClientInfo
}

func NewClientManager(proxyServer *tunnel.WebsocketHTTPProxyServer) *ClientManager {
	mngr := &ClientManager{
		proxyServer: proxyServer,
		clientsInfo: make(map[string]*ClientInfo),
	}
	proxyServer.SetConnectHandler(mngr.ClientConnected)
	proxyServer.SetDisconnectHandler(mngr.ClientDisconnected)
	return mngr
}

func (m *ClientManager) ClientConnected(sessionID string) {
	req, err := http.NewRequest(http.MethodGet, "/info", nil)
	if err != nil {
		log.Printf("error creating info request: %v", err)
		return
	}
	resp, err := m.proxyServer.ExecuteRequest(sessionID, req)
	if err != nil {
		log.Printf("info request error: %v", err)
		return
	}
	var info ClientInfo
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		log.Printf("error decoding client info: %v", err)
		return
	}
	// TODO: mutex
	m.clientsInfo[sessionID] = &info
	fmt.Println("got client info", info)
}

func (m *ClientManager) ClientDisconnected(sessionID string) {
	// TODO: mutex
	delete(m.clientsInfo, sessionID)
	fmt.Println("deleted client", sessionID)
}

func (m *ClientManager) Clients() map[string]*ClientInfo {
	res := make(map[string]*ClientInfo)

	// TODO: mutex
	for k, v := range m.clientsInfo {
		res[k] = v
	}

	return res
}

type ClientQueryResult struct {
	resp *http.Response
	err  error
}

func (m *ClientManager) QueryAll(r *http.Request) map[string]ClientQueryResult {
	res := make(map[string]ClientQueryResult)

	var wg sync.WaitGroup
	for sessionID := range m.clientsInfo {
		wg.Add(1)
		go func(sessionID string) {
			resp, err := m.proxyServer.ExecuteRequest(sessionID, r)
			// TODO: mutex
			res[sessionID] = ClientQueryResult{
				resp: resp,
				err:  err,
			}
			wg.Done()
		}(sessionID)
	}

	wg.Wait()

	return res
}
