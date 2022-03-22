package render

import (
	"github.com/techarm/go-bookings/internal/models"
	"net/http"
	"testing"
)

func TestNewTemplate(t *testing.T) {
	NewRenderer(app)
}

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	infoData := "info test data"
	session.Put(r.Context(), "info", infoData)
	result := AddDefaultData(r, &td)
	if result.Info != infoData {
		t.Errorf("expected %s, but got %s", infoData, result.Info)
	}
}

func TestExecute(t *testing.T) {
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	pathToTemplates = "./../../templates"
	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}

	app.UseCache = true
	app.TemplateCache = tc

	var w myWrite
	var td models.TemplateData

	err = Template(&w, r, "home.page.tmpl", &td)
	if err != nil {
		t.Error(err)
	}

	err = Template(&w, r, "not-exist.page.tmpl", &td)
	if err == nil {
		t.Error(err)
	}
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../templates"
	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}

	if len(tc) == 0 {
		t.Error("テンプレートキャッシュデータが生成できません")
	}
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		return nil, err
	}

	ctx := r.Context()
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	r = r.WithContext(ctx)
	return r, nil
}
