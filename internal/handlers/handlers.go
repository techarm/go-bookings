package handlers

import (
	"encoding/json"
	"github.com/techarm/go-bookings/internal/config"
	"github.com/techarm/go-bookings/internal/models"
	"github.com/techarm/go-bookings/internal/render"
	"log"
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

// Home Home画面の表示処理
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	stringMap := make(map[string]string)
	stringMap["message"] = "こんにちは！"
	stringMap["remote_ip"] = remoteIP

	render.Execute(w, r, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// About About画面の表示処理
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["remote_ip"] = m.App.Session.GetString(r.Context(), "remote_ip")
	render.Execute(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// GeneralsQuarters GeneralsQuarters画面の表示処理
func (m *Repository) GeneralsQuarters(w http.ResponseWriter, r *http.Request) {
	render.Execute(w, r, "rooms-generals.page.tmpl", &models.TemplateData{})
}

// MajorsSuite MajorsSuite画面の表示処理
func (m *Repository) MajorsSuite(w http.ResponseWriter, r *http.Request) {
	render.Execute(w, r, "rooms-majors.page.tmpl", &models.TemplateData{})
}

// SearchAvailability 予約状況検索画面の表示処理
func (m *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	render.Execute(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

// PostSearchAvailability 予約状況検索画面のPOST処理
func (m *Repository) PostSearchAvailability(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("post data from form"))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// SearchAvailabilityJSON 予約状況検索画面のAPI処理
func (m *Repository) SearchAvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	response := &jsonResponse{
		OK:      false,
		Message: "Available!",
	}

	jsonText, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		log.Println("JSONシリアライズ失敗", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonText)
}

// MakeReservation 予約画面の表示処理
func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	render.Execute(w, r, "make-reservation.page.tmpl", &models.TemplateData{})
}

// Contact 連絡画面の表示処理
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.Execute(w, r, "contact.page.tmpl", &models.TemplateData{})
}
