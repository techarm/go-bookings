package handlers

import (
	"encoding/gob"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/justinas/nosurf"
	"github.com/techarm/go-bookings/internal/config"
	"github.com/techarm/go-bookings/internal/helpers"
	"github.com/techarm/go-bookings/internal/models"
	"github.com/techarm/go-bookings/internal/render"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var app *config.AppConfig
var session *scs.SessionManager
var pathToTemplates = "./../../templates"
var functions = template.FuncMap{}

func getRouters() http.Handler {
	// セッション情報モデル登録
	gob.Register(models.Reservation{})

	app = &config.AppConfig{
		UseCache:     true,
		InProduction: false,
	}

	// テンプレートのキャッシュマップを作成
	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal("テンプレートキャッシュが作成できませんでした", err)
	}
	app.TemplateCache = tc

	// セッション管理設定
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	// setup logger
	infoLog := log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
	app.InfoLog = infoLog
	app.ErrorLog = errorLog

	render.NewRenderer(app)

	repo := NewRepository(app, nil)
	NewHandlers(repo)

	helpers.NewHelpers(app)

	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	//mux.Use(CSRFToken)
	mux.Use(SessionLoad)

	mux.Get("/", Repo.Home)
	mux.Get("/about", Repo.About)
	mux.Get("/rooms/majors-suite", Repo.MajorsSuite)
	mux.Get("/rooms/generals-quarters", Repo.GeneralsQuarters)

	mux.Get("/search-availability", Repo.SearchAvailability)
	mux.Post("/search-availability", Repo.PostSearchAvailability)
	mux.Post("/search-availability-json", Repo.SearchAvailabilityJSON)

	mux.Get("/make-reservation", Repo.MakeReservation)
	mux.Post("/make-reservation", Repo.PostMakeReservation)
	mux.Get("/reservation-summary", Repo.ReservationSummary)

	mux.Get("/contact", Repo.Contact)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}

// CSRFToken CSRF生成のミドルウェア
func CSRFToken(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad リクエストのセッション取得と保存処理を行う
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

// CreateTemplateCache テンプレートのキャッシュマップを作成
func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		t, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		match, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
		if err != nil {
			return cache, err
		}

		if len(match) > 0 {
			t, err = t.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
			if err != nil {
				return cache, err
			}
		}

		cache[name] = t
	}

	return cache, nil
}
