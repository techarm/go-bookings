package main

import (
	"encoding/gob"
	"github.com/alexedwards/scs/v2"
	"github.com/techarm/go-bookings/internal/config"
	"github.com/techarm/go-bookings/internal/driver"
	"github.com/techarm/go-bookings/internal/handlers"
	"github.com/techarm/go-bookings/internal/helpers"
	"github.com/techarm/go-bookings/internal/models"
	"github.com/techarm/go-bookings/internal/render"
	"github.com/techarm/go-bookings/internal/repository/dbrepo"
	"log"
	"net/http"
	"os"
	"time"
)

const port = ":8080"

var app *config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

// main アプリのメイン処理
func main() {
	db, err := run()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.SQL.Close()

	infoLog.Printf("Starting application on port %s\n", port)
	srv := &http.Server{
		Addr:    port,
		Handler: routers(app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		errorLog.Fatalln("サーバーが起動できませんでした。", err)
	}
}

// run アプリ起動処理
func run() (*driver.DB, error) {
	// セッション情報モデル登録
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Reservation{})
	gob.Register(models.Restriction{})
	gob.Register(models.RoomRestriction{})

	// setup logger
	infoLog = log.New(os.Stdout, "[INFO ] ", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)

	app = &config.AppConfig{
		UseCache:     false,
		InProduction: false,
		InfoLog:      infoLog,
		ErrorLog:     errorLog,
	}

	// テンプレートのキャッシュマップを作成
	tc, err := render.CreateTemplateCache()
	if err != nil {
		errorLog.Fatalln("テンプレートキャッシュが作成できませんでした", err)
		return nil, err
	}
	app.TemplateCache = tc

	// セッション管理設定
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	// set database
	infoLog.Println("connecting to database.")

	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=bookings user=techarm password=")
	if err != nil {
		errorLog.Fatalln("can not connect to database.")
	}

	infoLog.Println("database connection was successful. ")
	dbrepo.NewPostgresRepo(db.SQL, app)

	render.NewRenderer(app)

	repo := handlers.NewRepository(app, db)
	handlers.NewHandlers(repo)

	helpers.NewHelpers(app)

	return db, nil
}
