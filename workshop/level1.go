package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// colors returns a list of 25 colors
func colors() []string {
	// step 1: return an array of 25x the same color (loop, arrays)
	// step 2: alternate between 2 colors (conditions)
	// step 3: make it a gradient (string interpolation)
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
