package handlers

import (
	"github.com/techarm/go-bookings/pkg/config"
	"github.com/techarm/go-bookings/pkg/models"
	"github.com/techarm/go-bookings/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// NewRepository リポジトリを作成する
func NewRepository(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers 初期化処理
func NewHandlers(r *Repository) {
	Repo = r
}

// Home 画面ハンドラー
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	stringMap := make(map[string]string)
	stringMap["message"] = "こんにちは！"
	stringMap["remote_ip"] = remoteIP

	render.Execute(w, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// About 画面ハンドラー
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["remote_ip"] = m.App.Session.GetString(r.Context(), "remote_ip")
	render.Execute(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
