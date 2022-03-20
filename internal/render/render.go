package render

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/justinas/nosurf"
	"github.com/techarm/go-bookings/internal/config"
	"github.com/techarm/go-bookings/internal/models"
	"html/template"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app *config.AppConfig
var pathToTemplates = "./templates"

// NewTemplate 設置初期化
func NewTemplate(a *config.AppConfig) {
	app = a
}

// AddDefaultData ディフォルトデータを設定する
func AddDefaultData(r *http.Request, td *models.TemplateData) *models.TemplateData {
	td.Info = app.Session.PopString(r.Context(), "info")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.CSRFToken = nosurf.Token(r)
	return td
}

// Execute html/templateを使い、テンプレートファイルをレンダリング
func Execute(w http.ResponseWriter, r *http.Request, name string, td *models.TemplateData) error {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[name]
	if !ok {
		//log.Fatalln("テンプレートキャッシュに対象データが存在しません: ", name)
		return errors.New(fmt.Sprintf("テンプレートキャッシュに対象データが存在しません: %s", name))
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(r, td)
	err := t.Execute(buf, td)
	if err != nil {
		//log.Println("テンプレートのbuffer書き込み失敗しました: ", err)
		return err
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		//log.Println("テンプレートの応答データ書き込み失敗しました: ", err)
		return err
	}

	return nil
}

// CreateTemplateCache テンプレートのキャッシュマップを作成
func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		t, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		match, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
		if err != nil {
			return cache, err
		}

		if len(match) > 0 {
			t, err = t.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
			if err != nil {
				return cache, err
			}
		}

		cache[name] = t
	}

	return cache, nil
}
