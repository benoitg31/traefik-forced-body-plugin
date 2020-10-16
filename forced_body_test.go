package plugin_forcedbody

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestServeHTTP(t *testing.T) {
	tests := []struct {
		desc            string
		contentEncoding string
		resBody         string
		expResBody      string
		expLastModified bool
	}{
		{
			desc:       "should replace foo by bar",
			expResBody: "forced body",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			config := &Config{
				Body: "forced body",
			}

			next := func(rw http.ResponseWriter, req *http.Request) {
				rw.Header().Set("Content-Encoding", test.contentEncoding)
				rw.Header().Set("Last-Modified", "Thu, 02 Jun 2016 06:01:08 GMT")
				rw.Header().Set("Content-Length", strconv.Itoa(len(test.resBody)))
				rw.WriteHeader(http.StatusOK)

				_, _ = fmt.Fprintf(rw, test.resBody)
			}

			rewriteBody, err := New(context.Background(), http.HandlerFunc(next), config, "rewriteBody")
			if err != nil {
				t.Fatal(err)
			}

			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/", nil)

			rewriteBody.ServeHTTP(recorder, req)

			if !bytes.Equal([]byte(test.expResBody), recorder.Body.Bytes()) {
				t.Errorf("got body %q, want %q", recorder.Body.Bytes(), test.expResBody)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		desc   string
		expErr bool
	}{
		{
			desc:   "should return no error",
			expErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			config := &Config{
				Body: "forced body",
			}

			_, err := New(context.Background(), nil, config, "rewriteBody")
			if test.expErr && err == nil {
				t.Fatal("expected error on bad regexp format")
			}
		})
	}
}
