package handler

import (
	"github.com/techarm/go-bookings/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.Render(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.Render(w, "about.page.html")
}
