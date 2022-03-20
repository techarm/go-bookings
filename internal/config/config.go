package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
	"log"
)

// AppConfig アプリケーションの全体設定
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InProduction  bool
	Session       *scs.SessionManager
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
}
