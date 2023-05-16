package main

import (
	"encoding/json"
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

	q := r.URL.Query().Get("query")

	gif, err := giphyGif(q)
	if err != nil {
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(struct {
		GifURL string `json:"gif_url"`
	}{GifURL: gif})
}

func giphyGif(search string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, "https://api.giphy.com/v1/videos/search", nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	urlValues := req.URL.Query()
	urlValues.Add("q", search)
	urlValues.Add("limit", "10")
	urlValues.Add("api_key", "Gc7131jiJuvI7IdN0HZ1D7nh0ow5BU6g")
	req.URL.RawQuery = urlValues.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("request error: %v", err)
	}
	defer resp.Body.Close()

	/*
		{
			"data": [
				"featured_gif": {
					"images": {
						"preview_gif": {
							"url": "..."
						}
					}
				}
			]
		}
	*/
	var giphyResult struct {
		Data []struct {
			EmbedURL string `json:"embed_url"`
			Images   struct {
				FixedWith struct {
					WebP string
				} `json:"fixed_width"`
			}
		}
	}

	err = json.NewDecoder(resp.Body).Decode(&giphyResult)
	if err != nil {
		return "", fmt.Errorf("JSON error: %w", err)
	}

	return giphyResult.Data[0].Images.FixedWith.WebP, nil
}
