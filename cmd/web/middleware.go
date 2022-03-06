package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)

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
