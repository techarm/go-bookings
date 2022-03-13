package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/techarm/go-bookings/internal/config"
	"github.com/techarm/go-bookings/internal/handlers"
	"net/http"
)

// routers ルーター設定処理
func routers(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(CSRFToken)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/rooms/majors-suite", handlers.Repo.MajorsSuite)
	mux.Get("/rooms/generals-quarters", handlers.Repo.GeneralsQuarters)

	mux.Get("/search-availability", handlers.Repo.SearchAvailability)
	mux.Post("/search-availability", handlers.Repo.PostSearchAvailability)
	mux.Post("/search-availability-json", handlers.Repo.SearchAvailabilityJSON)

	mux.Get("/make-reservation", handlers.Repo.MakeReservation)
	mux.Get("/contact", handlers.Repo.Contact)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
