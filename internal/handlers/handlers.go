package handlers

import (
	"encoding/json"
	"github.com/techarm/go-bookings/internal/config"
	"github.com/techarm/go-bookings/internal/driver"
	"github.com/techarm/go-bookings/internal/forms"
	"github.com/techarm/go-bookings/internal/helpers"
	"github.com/techarm/go-bookings/internal/models"
	"github.com/techarm/go-bookings/internal/render"
	"github.com/techarm/go-bookings/internal/repository"
	"github.com/techarm/go-bookings/internal/repository/dbrepo"
	"net/http"
	"strconv"
	"time"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// NewRepository リポジトリを作成する
func NewRepository(a *config.AppConfig, d *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(d.SQL, a),
	}
}

// NewHandlers 初期化処理
func NewHandlers(r *Repository) {
	Repo = r
}

// Home Home画面の表示処理
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About About画面の表示処理
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "about.page.tmpl", &models.TemplateData{})
}

// GeneralsQuarters GeneralsQuarters画面の表示処理
func (m *Repository) GeneralsQuarters(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "rooms-generals.page.tmpl", &models.TemplateData{})
}

// MajorsSuite MajorsSuite画面の表示処理
func (m *Repository) MajorsSuite(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "rooms-majors.page.tmpl", &models.TemplateData{})
}

// SearchAvailability 予約状況検索画面の表示処理
func (m *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "search-availability.page.tmpl", &models.TemplateData{})
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
		helpers.ServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonText)
}

// MakeReservation 予約画面の表示処理
func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	var emptyReservation = models.Reservation{
		StartDate: time.Now(),
		EndDate:   time.Now().Add(24 * time.Hour),
		RoomID:    1,
	}

	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	render.Template(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// PostMakeReservation 予約画面の登録処理
func (m *Repository) PostMakeReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	var dateLayout = "2006/01/02"
	startDate, err := time.Parse(dateLayout, r.Form.Get("start_date"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	endDate, err := time.Parse(dateLayout, r.Form.Get("start_date"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	roomID, err := strconv.Atoi(r.Form.Get("room_id"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
		StartDate: startDate,
		EndDate:   endDate,
		RoomID:    roomID,
	}

	form := forms.New(r.PostForm)
	form.Required("first_name", "last_name", "email", "phone")
	form.MinLength("first_name", 2)
	form.MinLength("last_name", 2)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation
		render.Template(w, r, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		m.App.ErrorLog.Printf("Form検証エラー: %v\n", form.Errors)
		return
	}

	// 予約データを保存する
	newReservationID, err := m.DB.InsertReservation(reservation)
	if err != nil {
		helpers.ServerError(w, err)
	}

	// 予約制限データを保存する
	var roomRestriction = models.RoomRestriction{
		StartDate:     startDate,
		EndDate:       endDate,
		RoomID:        roomID,
		ReservationID: newReservationID,
		RestrictionID: 1,
	}

	err = m.DB.InsertRoomRestriction(roomRestriction)
	if err != nil {
		helpers.ServerError(w, err)
	}

	m.App.InfoLog.Printf("予約データを登録しました。予約ID:%d\n", newReservationID)

	m.App.Session.Put(r.Context(), "reservation", reservation)
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

// ReservationSummary 予約内容確認画面表示処理
func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		m.App.ErrorLog.Println("セッションから予約情報が取得できません")
		m.App.Session.Put(r.Context(), "error", "セッション情報取得失敗しました、最初からやり直ししてください。")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	m.App.Session.Remove(r.Context(), "reservation")

	data := make(map[string]interface{})
	data["reservation"] = reservation

	render.Template(w, r, "reservation-summary.page.tmpl", &models.TemplateData{
		Data: data,
	})
}

// Contact 連絡画面の表示処理
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "contact.page.tmpl", &models.TemplateData{})
}
