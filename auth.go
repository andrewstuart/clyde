package clyde

import "net/http"

type Auther interface {
	AuthHeader() string
}

type AuthRoundTripper struct {
	RT     http.RoundTripper
	Auther Auther
}

func (a *AuthRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if a.RT == nil {
		a.RT = http.DefaultTransport
	}

	req.Header.Set("Authorization", a.Auther.AuthHeader())
	return a.RT.RoundTrip(req)
}
