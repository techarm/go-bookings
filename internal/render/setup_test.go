package render

import (
	"encoding/gob"
	"github.com/alexedwards/scs/v2"
	"github.com/techarm/go-bookings/internal/config"
	"github.com/techarm/go-bookings/internal/models"
	"net/http"
	"os"
	"testing"
	"time"
)

var testApp config.AppConfig
var session *scs.SessionManager

func TestMain(m *testing.M) {
	// テスト実行前の環境設定
	gob.Register(models.Reservation{})

	// セッション管理設定
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false
	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())
}

type myWrite struct{}

func (mw *myWrite) Header() http.Header {
	var h http.Header
	return h
}

func (mw *myWrite) Write(b []byte) (int, error) {
	return len(b), nil
}

func (mw *myWrite) WriteHeader(statusCode int) {
}
