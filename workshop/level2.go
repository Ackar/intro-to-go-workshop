package main

import (
	"encoding/json"
	"net/http"
)

type Level2 struct {
	idx int
}

func NewLevel2() *Level2 {
	return &Level2{
		// FIXME
	}
}

func (l *Level2) colors() []string {
	// step 1: make a square move at every invocation (store state)
	// step 2: use defer to change the state
	// step 3: get the color as a constructor parameter
	var res []string
	for i := 0; i < 25; i++ {
		if i == l.idx {
			res = append(res, "silver")
		} else {
			res = append(res, "blue")
		}
	}
	l.idx = (l.idx + 1) % 25
	return res
}

func (l *Level2) Handler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(colorsResponse{
		Colors: l.colors(),
	})
}
