package main

import (
	"encoding/json"
	"net/http"
)

// Level2 is about using structs and storing state.
// Use the struct below to store a state and make colors move!
//
// Step 1: make a square move at every invocation (learn about structs)
// Step 2: use defer to change the state (learn about defer)
// Step 3: get the color as a constructor parameter (learn how to add parameters to a struct)
type Level2 struct {
	// FIXME
}

func NewLevel2() *Level2 {
	return &Level2{
		// FIXME
	}
}

func (l *Level2) colors() []string {
	// FIXME
	return nil
}

// No need to edit below this line

func (l *Level2) Handler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(colorsResponse{
		Colors: l.colors(),
	})
}
