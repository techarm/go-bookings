package main

import (
	"github.com/alexedwards/scs/v2"
	"github.com/techarm/go-bookings/pkg/config"
	"github.com/techarm/go-bookings/pkg/handlers"
	"github.com/techarm/go-bookings/pkg/render"
	"log"
	"net/http"
	"time"
)

const port = ":8080"

var app *config.AppConfig
var session *scs.SessionManager

func main() {

	app = &config.AppConfig{
		UseCache:     false,
		InProduction: false,
	}

	// テンプレートのキャッシュマップを作成
	tc, err := render.CreateTemplateCache()
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

	render.NewTemplate(app)

	r := handlers.NewRepository(app)
	handlers.NewHandlers(r)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	log.Println("start server and listen on", port)
	srv := &http.Server{
		Addr:    port,
		Handler: routers(app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalln("サーバーが起動できませんでした。", err)
	}
}
