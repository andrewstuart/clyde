package clyde

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeaderRTer(t *testing.T) {
	asrt := assert.New(t)

	fb := HeaderRoundTripper(map[string]string{
		"foo":        "foo",
		"bar":        "baz",
		"banterlasd": "wefjasodijfews",
		"Assertion":  "YesItsHere",
	})

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		for k, v := range fb {
			asrt.Equal(r.Header.Get(k), v)
		}

		fmt.Fprintf(w, "hello")
	}))
	defer ts.Close()

	c := &http.Client{Transport: fb}

	res, err := c.Get(ts.URL)
	asrt.NoError(err)

	body, err := ioutil.ReadAll(res.Body)
	asrt.NoError(err)
	asrt.Equal("hello", string(body))
}
