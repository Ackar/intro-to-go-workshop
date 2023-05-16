package main

import (
	"fmt"
	"net/http"
)

// Level3Handler returns a link to your a GIF.
// It should return a JSON in the following format:
//
//	{
//		 "gif_url": "URL"
//	}
//
// Step 1: return a static JSON containing a link to your favorite GIF
// Step 2: fetch a GIF from Giphy and return it
// Step 3: get the Giphy search string from the "query" query parameter
func Level3Handler(w http.ResponseWriter, r *http.Request) {
	// FIXME
}

// giphyGif returns the first GIF returns by the given Giphy search
func giphyGif(search string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, "https://api.giphy.com/v1/videos/search", nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	urlValues := req.URL.Query()
	urlValues.Add("q", search)
	urlValues.Add("limit", "10")
	urlValues.Add("api_key", "FIXME")
	req.URL.RawQuery = urlValues.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("request error: %v", err)
	}
	defer resp.Body.Close()

	// The result is in the following format (relevant part only):
	// {
	// 	"data": [
	// 		"images": {
	// 			"fixed_width": {
	// 				"webp": "URL"
	// 			}
	// 		}
	// 	]
	// }

	// FIXME

	return "", fmt.Errorf("unimplemented")
}
