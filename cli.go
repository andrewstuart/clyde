package clyde

import "net/http"

type RoundTripFunc func(*http.Request) (*http.Response, error)

func (r RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return r(req)
}

type HeaderRoundTripper map[string]string

func (h HeaderRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	for k, v := range h {
		req.Header.Set(k, v)
	}

	return http.DefaultTransport.RoundTrip(req)
}
