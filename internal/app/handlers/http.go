package handlers

import (
	"fmt"
	"net/http"
)

func (p *Proxy) HandleHTTP(req *http.Request) (*http.Response, error) {
	request, err := http.NewRequest(req.Method, req.RequestURI, req.Body)
	if err != nil {
		return nil, fmt.Errorf("could not create new request: %w", err)
	}

	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	for key, values := range req.Header {
		for _, value := range values {
			request.Header.Add(key, value)
		}
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("сould not execute request: %w", err)
	}

	return resp, nil
}
