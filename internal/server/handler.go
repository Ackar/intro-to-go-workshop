package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"sync"
	"time"
)

type Resolver struct {
	manager *ClientManager

	level1Mutex       sync.Mutex
	level1Cache       []clientColors
	level1RefreshedAt time.Time

	level2Mutex       sync.Mutex
	level2Cache       []clientColors
	level2RefreshedAt time.Time

	level3Mutex       sync.Mutex
	level3Cache       []clientGIF
	level3RefreshedAt time.Time
}

func NewLevelHandler(clientManager *ClientManager) *Resolver {
	return &Resolver{
		manager: clientManager,
	}
}

func (h *Resolver) Clients(ctx context.Context) []ClientInfo {
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

func (h *Resolver) Level1(ctx context.Context) ([]clientColors, error) {
	h.level1Mutex.Lock()
	defer h.level1Mutex.Unlock()
	if time.Since(h.level1RefreshedAt) < 3*time.Second {
		return h.level1Cache, nil
	}

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

	sort.Slice(res, func(i, j int) bool {
		return res[i].Client.Name < res[j].Client.Name
	})

	h.level1Cache = res
	h.level1RefreshedAt = time.Now()

	return res, nil
}

func (h *Resolver) Level2(ctx context.Context) ([]clientColors, error) {
	h.level2Mutex.Lock()
	defer h.level2Mutex.Unlock()
	if time.Since(h.level2RefreshedAt) < 3*time.Second {
		return h.level2Cache, nil
	}

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

	sort.Slice(res, func(i, j int) bool {
		return res[i].Client.Name < res[j].Client.Name
	})

	h.level2Cache = res
	h.level2RefreshedAt = time.Now()

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

func (h *Resolver) Level3(ctx context.Context) ([]clientGIF, error) {
	h.level3Mutex.Lock()
	defer h.level3Mutex.Unlock()
	if time.Since(h.level3RefreshedAt) < 3*time.Second {
		return h.level3Cache, nil
	}

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

	sort.Slice(res, func(i, j int) bool {
		return res[i].Client.Name < res[j].Client.Name
	})

	h.level3Cache = res
	h.level3RefreshedAt = time.Now()

	return res, nil
}
