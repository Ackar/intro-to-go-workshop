package main

import (
	"encoding/json"
	"fmt"
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
	var res []string
	color := "rgb(161, %d, 235)"
	for i := 0; i < 25; i++ {
		if i%2 == 0 {
			res = append(res, fmt.Sprintf(color, 100-i*2))
		} else {
			res = append(res, fmt.Sprintf(color, 100+i*4))
		}
	}
	return res
}

type colorsResponse struct {
	Colors []string `json:"colors"`
}

func Level1Handler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(colorsResponse{
		Colors: colors(),
	})
}
