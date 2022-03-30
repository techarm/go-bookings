package helpers

import (
	"github.com/techarm/go-bookings/internal/config"
	"net/http"
	"runtime/debug"
)

var app *config.AppConfig

func NewHelpers(a *config.AppConfig) {
	app = a
}

func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Printf("クライアントエラーが発生しました。statusCode=%d\n", status)
	http.Error(w, http.StatusText(status), status)
}

func SystemErrorAndRedirectToRoot(w http.ResponseWriter, r *http.Request, errs ...error) {
	if len(errs) > 0 {
		app.ErrorLog.Println(errs[0])
	}
	app.Session.Put(r.Context(), "error", "システムエラーが発生しました。ブラウザを閉じて再度やり直してください。")
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func ServerError(w http.ResponseWriter, err error) {
	app.ErrorLog.Printf("%s\n%s", err.Error(), debug.Stack())
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
