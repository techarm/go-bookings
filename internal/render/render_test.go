package render

import (
	"github.com/techarm/go-bookings/internal/models"
	"net/http"
	"testing"
)

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
