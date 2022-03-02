package config

import "html/template"

// AppConfig アプリケーションの全体設定
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
}
