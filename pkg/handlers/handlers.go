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
	stringMap := make(map[string]string)
	stringMap["message"] = "こんにちは！"

	render.Execute(w, "home.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

// About 画面ハンドラー
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.Execute(w, "about.page.html", &models.TemplateData{})
}
