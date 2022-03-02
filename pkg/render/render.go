package render

import (
	"bytes"
	"github.com/techarm/go-bookings/pkg/config"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplate 設置初期化
func NewTemplate(a *config.AppConfig) {
	app = a
}

// Execute html/templateを使い、テンプレートファイルをレンダリング
func Execute(w http.ResponseWriter, name string) {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[name]
	if !ok {
		log.Fatalln("テンプレートキャッシュに対象データが存在しません", name)
	}

	buf := new(bytes.Buffer)
	err := t.Execute(buf, nil)
	if err != nil {
		log.Fatalln("テンプレートのbuffer書き込み失敗しました", err)
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Fatalln("テンプレートの応答データ書き込み失敗しました", err)
	}
}

// CreateTemplateCache テンプレートのキャッシュマップを作成
func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		t, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		match, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return cache, err
		}

		if len(match) > 0 {
			t, err = t.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return cache, err
			}
		}

		cache[name] = t
	}

	return cache, nil
}
