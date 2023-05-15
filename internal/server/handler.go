package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type LevelHandler struct {
	manager *ClientManager
}

func NewLevelHandler(clientManager *ClientManager) *LevelHandler {
	return &LevelHandler{
		manager: clientManager,
	}
}

func (h *LevelHandler) Clients(ctx context.Context) []ClientInfo {
	clients := h.manager.Clients()

	var res []ClientInfo
	for _, cl := range clients {
		res = append(res, *cl)
	}

	return res
}

type colorsResponse struct {
	Colors []string `json:"colors"`
}

type clientColors struct {
	Client *ClientInfo `json:"client_info"`
	Colors []string    `json:"squares"`
	Error  *string     `json:"error"`
}

func (h *LevelHandler) Level1(ctx context.Context) ([]clientColors, error) {
	req, err := http.NewRequest(http.MethodGet, "/level1", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request")
	}

	results := h.manager.QueryAll(req)
	clients := h.manager.Clients()

	var res []clientColors

	for sessID, r := range results {
		client, ok := clients[sessID]
		if !ok {
			continue
		}
		var clientResult colorsResponse
		if r.err == nil {
			r.err = json.NewDecoder(r.resp.Body).Decode(&clientResult)
		}
		var errPtr *string
		if r.err != nil {
			errStr := r.err.Error()
			errPtr = &errStr
		}
		res = append(res, clientColors{
			Client: client,
			Colors: clientResult.Colors,
			Error:  errPtr,
		})
	}

	return res, nil
}

func (h *LevelHandler) Level2(ctx context.Context) ([]clientColors, error) {
	req, err := http.NewRequest(http.MethodGet, "/level2", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request")
	}

	results := h.manager.QueryAll(req)
	clients := h.manager.Clients()

	var res []clientColors

	for sessID, r := range results {
		client, ok := clients[sessID]
		if !ok {
			continue
		}
		var clientResult colorsResponse
		if r.err == nil {
			r.err = json.NewDecoder(r.resp.Body).Decode(&clientResult)
		}
		var errPtr *string
		if r.err != nil {
			errStr := r.err.Error()
			errPtr = &errStr
		}
		res = append(res, clientColors{
			Client: client,
			Colors: clientResult.Colors,
			Error:  errPtr,
		})
	}

	return res, nil
}

type gifResponse struct {
	GIFUrl string `json:"gif_url"`
}

type clientGIF struct {
	Client *ClientInfo
	GIFUrl string
	Error  *string
}

func (h *LevelHandler) Level3(ctx context.Context) ([]clientGIF, error) {
	req, err := http.NewRequest(http.MethodGet, "/level3?query=gopher", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request")
	}

	results := h.manager.QueryAll(req)
	clients := h.manager.Clients()

	var res []clientGIF

	for sessID, r := range results {
		client, ok := clients[sessID]
		if !ok {
			continue
		}
		var clientResult gifResponse
		if r.err == nil {
			r.err = json.NewDecoder(r.resp.Body).Decode(&clientResult)
		}
		var errPtr *string
		if r.err != nil {
			errStr := r.err.Error()
			errPtr = &errStr
		}
		res = append(res, clientGIF{
			Client: client,
			GIFUrl: clientResult.GIFUrl,
			Error:  errPtr,
		})
	}

	return res, nil
}
