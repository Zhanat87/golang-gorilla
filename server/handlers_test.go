/*
https://golang.org/pkg/net/http/httptest/#example_ResponseRecorder
https://elithrar.github.io/article/testing-http-handlers-go/
 */
package server

import (
	"net/http"
	"testing"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Ping)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "pong", rr.Body.String())
}
