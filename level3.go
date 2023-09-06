package main

import (
	"fmt"
	"net/http"
)

// Level3Handler is a HTTP handler that returns a link to a GIF in the JSON
// format.
// The returned JSON should have the following format:
//
//	{
//		 "gif_url": "URL"
//	}
//
// Step 1: return a static JSON containing a link to your favorite GIF. Use `w`
// to send data back to the client and the `json` package to format your JSON.
// See https://pkg.go.dev/encoding/json (hint: look at NewEncoder or Marshal).\
//
// Step 2: fetch a GIF from Giphy and return it. See the gifURL function below.
//
// Step 3: Get the "query" query parameter from the HTTP request and use it to
// call the gifURL function.
// The request uses the following format:
//
//	/level3?query=search
//
// This means you can get the query and use it in your search!
// The http.Request parameter contains information about the current HTTP
// request, look into r.URL to find the parameter!
func Level3Handler(w http.ResponseWriter, r *http.Request) {
	// FIXME
}

// Step 2/3 only
// gifURL returns the first GIF returned by the given Giphy search
func gifURL(search string) (string, error) {
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
	//		{
	// 			"images": {
	// 				"fixed_width": {
	// 					"webp": "URL"
	// 				}
	// 			}
	//		}
	// 	]
	// }

	// FIXME

	return "", fmt.Errorf("unimplemented")
}
