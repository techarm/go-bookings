package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestCSRFToken(t *testing.T) {
	var th testHandler
	h := CSRFToken(&th)

	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Error(fmt.Sprintf("type is not http.Handler, but is %T", v))
	}
}

func TestSessionLoad(t *testing.T) {
	var th testHandler
	h := SessionLoad(&th)

	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Error(fmt.Sprintf("type is not http.Handler, but is %T", v))
	}
}
