package main

import (
	"encoding/json"
	"net/http"
)

// colors returns a list of 25 colors. The colors can be any string representing
// a CSS color.
// Some of the possible formats: "blue", "rgb(42, 42, 42)", // "#FF00FF".
//
// Step 1: return an array of 25x the same color (learn about loops and arrays)
// Step 2: alternate between 2 colors (learn about conditions)
// Step 3: make it a gradient (learn to build strings)
func colors() []string {
	// FIXME
	return nil
}

// No need to edit below this line

type colorsResponse struct {
	Colors []string `json:"colors"`
}

func Level1Handler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(colorsResponse{
		Colors: colors(),
	})
}
