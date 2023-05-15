package tunnel

import (
	"bytes"
	"io"
	"net/http"

	"github.com/google/uuid"
)

type HTTPRequest struct {
	ID     string
	Method string
	URL    string
	Header http.Header
	Body   []byte
}

func FromHTTPRequest(r *http.Request) *HTTPRequest {
	var body []byte
	if r.Body != nil {
		body, _ = io.ReadAll(r.Body)
	}

	return &HTTPRequest{
		ID:     uuid.New().String(),
		Method: r.Method,
		URL:    r.URL.String(),
		Header: r.Header,
		Body:   body,
	}
}

type HTTPResponse struct {
	ID         string
	StatusCode int
	Header     http.Header
	Body       []byte
}

func (h *HTTPResponse) ToResponse() *http.Response {
	return &http.Response{
		StatusCode: h.StatusCode,
		Header:     h.Header,
		Body:       io.NopCloser(bytes.NewReader(h.Body)),
	}
}

func FromHTTPResponse(id string, resp *http.Response) *HTTPResponse {
	body, _ := io.ReadAll(resp.Body)
	return &HTTPResponse{
		ID:         id,
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		Body:       body,
	}
}
