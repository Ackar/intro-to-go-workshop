package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"sync"
	"time"
)

type Resolver struct {
	manager *ClientManager

	level1Mutex sync.Mutex
	level2Mutex sync.Mutex
	level3Mutex sync.Mutex

	// cache fields to avoid requesting clients often when we have multiple
	// users on the website
	level1Cache fieldCache[[]clientColors]
	level2Cache fieldCache[[]clientColors]
	level3Cache fieldCache[[]clientGIF]
}

func NewLevelHandler(clientManager *ClientManager) *Resolver {
	return &Resolver{
		manager:     clientManager,
		level1Cache: fieldCache[[]clientColors]{},
		level2Cache: fieldCache[[]clientColors]{},
		level3Cache: fieldCache[[]clientGIF]{},
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

func (r *colorsResponse) Validate() error {
	if len(r.Colors) != 25 {
		return fmt.Errorf("invalid length: %d", len(r.Colors))
	}
	return nil
}

type clientColors struct {
	Client *ClientInfo `json:"client_info"`
	Colors []string    `json:"squares"`
	Error  *string     `json:"error"`
}

func (h *Resolver) Level1(ctx context.Context) ([]clientColors, error) {
	h.level1Mutex.Lock()
	defer h.level1Mutex.Unlock()

	if cachedRes, valid := h.level1Cache.Get(); valid {
		return cachedRes, nil
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
			if r.err == nil {
				r.err = clientResult.Validate()
			}
		}
		var errPtr *string
		if r.err != nil {
			errStr := r.err.Error()
			errPtr = &errStr
			clientResult.Colors = nil
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

	h.level1Cache.Set(res)

	return res, nil
}

func (h *Resolver) Level2(ctx context.Context) ([]clientColors, error) {
	h.level2Mutex.Lock()
	defer h.level2Mutex.Unlock()

	if cachedRes, valid := h.level2Cache.Get(); valid {
		return cachedRes, nil
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
			if r.err == nil {
				r.err = clientResult.Validate()
			}
		}
		var errPtr *string
		if r.err != nil {
			errStr := r.err.Error()
			errPtr = &errStr
			clientResult.Colors = nil
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

	h.level2Cache.Set(res)

	return res, nil
}

type gifResponse struct {
	GIFUrl string `json:"gif_url"`
}

func (r *gifResponse) Validate() error {
	_, err := url.Parse(r.GIFUrl)
	if err != nil {
		return fmt.Errorf("invalid URL %q", r.GIFUrl)
	}
	return nil
}

type clientGIF struct {
	Client *ClientInfo
	GIFUrl string
	Error  *string
}

func (h *Resolver) Level3(ctx context.Context) ([]clientGIF, error) {
	h.level3Mutex.Lock()
	defer h.level3Mutex.Unlock()
	if cachedRes, valid := h.level3Cache.Get(); valid {
		return cachedRes, nil
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
			if r.err == nil {
				r.err = clientResult.Validate()
			}
		}
		var errPtr *string
		if r.err != nil {
			errStr := r.err.Error()
			errPtr = &errStr
			clientResult.GIFUrl = ""
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

	h.level3Cache.Set(res)

	return res, nil
}

type fieldCache[T any] struct {
	data        T
	refreshedAt time.Time
}

func (f *fieldCache[T]) Set(d T) {
	f.data = d
	f.refreshedAt = time.Now()
}

func (f *fieldCache[T]) Get() (T, bool) {
	if time.Since(f.refreshedAt) < 3*time.Second {
		return f.data, true
	}
	var res T
	return res, false
}
