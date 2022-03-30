package handlers

import (
	"context"
	"github.com/techarm/go-bookings/internal/models"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var testCases = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	//{"home", "/", "GET", []postData{}, http.StatusOK},
	//{"about", "/about", "GET", []postData{}, http.StatusOK},
	//{"majors-suite", "/rooms/majors-suite", "GET", []postData{}, http.StatusOK},
	//{"generals-quarters", "/rooms/generals-quarters", "GET", []postData{}, http.StatusOK},
	//{"make-reservation", "/make-reservation", "GET", []postData{}, http.StatusOK},
	//{"search-availability", "/search-availability", "GET", []postData{}, http.StatusOK},
	////{"reservation-summary", "/reservation-summary", "GET", []postData{}, http.StatusOK},
	//{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	//
	//{"search-availability-post", "/search-availability", "POST", []postData{
	//	{key: "start", value: "2022/03/19"},
	//	{key: "end", value: "2022/03/20"},
	//}, http.StatusOK},
	//{"search-availability-json", "/search-availability-json", "POST", []postData{
	//	{key: "start", value: "2022/03/19"},
	//	{key: "end", value: "2022/03/20"},
	//}, http.StatusOK},
	//{"make-reservation-post", "/make-reservation", "POST", []postData{
	//	{key: "user_name", value: "富田太郎"},
	//	{key: "email", value: "kimura@techarm.com"},
	//	{key: "phone_number", value: "08012345678"},
	//}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routers := getRouters()

	ts := httptest.NewTLSServer(routers)
	defer ts.Close()

	for _, e := range testCases {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		} else {
			values := url.Values{}
			for _, x := range e.params {
				values.Add(x.key, x.value)
			}

			resp, err := ts.Client().PostForm(ts.URL+e.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}

func TestRepository_MakeReservation(t *testing.T) {
	testCases := []struct {
		name       string
		data       *models.Reservation
		expectCode int
	}{
		{
			name:       "セッション取得失敗",
			data:       nil,
			expectCode: http.StatusTemporaryRedirect,
		},
		{
			name: "セッション取得正常",
			data: &models.Reservation{
				RoomID: 1,
			},
			expectCode: http.StatusOK,
		},
		{
			name: "RoomID取得失敗",
			data: &models.Reservation{
				RoomID: 100,
			},
			expectCode: http.StatusTemporaryRedirect,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/make-reservation", nil)
			ctx := getContext(req)
			req = req.WithContext(ctx)

			rr := httptest.NewRecorder()
			if tc.data != nil {
				session.Put(ctx, "reservation", *tc.data)
			}

			handler := http.HandlerFunc(Repo.MakeReservation)
			handler.ServeHTTP(rr, req)

			if rr.Code != tc.expectCode {
				t.Errorf("http status code got %d, but wanted %d", rr.Code, tc.expectCode)
			}
		})
	}
}

func getContext(r *http.Request) context.Context {
	ctx, err := session.Load(r.Context(), r.Header.Get("X-Session"))
	if err != nil {
		log.Println(err)
	}
	return ctx
}
