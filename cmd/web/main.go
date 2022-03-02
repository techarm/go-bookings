package main

import (
	"github.com/techarm/go-bookings/pkg/config"
	"github.com/techarm/go-bookings/pkg/handlers"
	"github.com/techarm/go-bookings/pkg/render"
	"log"
	"net/http"
)

const port = ":8080"

func main() {

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("テンプレートキャッシュが作成できませんでした", err)
	}

	app := &config.AppConfig{
		UseCache:      false,
		TemplateCache: tc,
	}

	render.NewTemplate(app)

	r := handlers.NewRepository(app)
	handlers.NewHandlers(r)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	log.Println("start server and listen on", port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("サーバーが起動できませんでした。", err)
	}
}
