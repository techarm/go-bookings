package handlers

import (
	"github.com/techarm/go-bookings/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.NewTemplate(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.NewTemplate(w, "about.page.html")
}
