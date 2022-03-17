package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/techarm/go-bookings/internal/config"
	"testing"
)

func TestRouters(t *testing.T) {
	var app config.AppConfig
	mux := routers(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing
	default:
		t.Error(fmt.Sprintf("type is not *chi.Mux, but is %T", v))
	}
}
