package clyde

import "net/http"

// RoundTripFunc creates an http.RoundTripper from a function that returns a
// response/error
type RoundTripFunc func(*http.Request) (*http.Response, error)

// RoundTrip implements http.RoundTripper
func (r RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return r(req)
}

// A HeaderRoundTripper is a map of header keys to header values. The
// HeaderRoundTripper is not safe for concurrent modification and use, as it is
// purely a raw map.
type HeaderRoundTripper map[string]string

// RoundTrip implements http.RoundTripper
func (h HeaderRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	for k, v := range h {
		req.Header.Set(k, v)
	}

	return http.DefaultTransport.RoundTrip(req)
}
