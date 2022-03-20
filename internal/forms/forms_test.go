package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("Form検証OKを期待していますが、検証NGになっています")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)
	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("必須項目が入力していないのに、Formの検証結果がOKになっています")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	r = httptest.NewRequest("POST", "/whatever", nil)
	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")

	if !form.Valid() {
		t.Error("必須項目が入力しているのに、Formの検証結果がNGになっています")
	}
}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	if form.Has("some-field") {
		t.Error("存在しない項目はfalseを想定しています")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "")
	form = New(postedData)

	if !form.Has("a") {
		t.Error("存在項目はtrueを想定しています")
	}

	if form.Has("b") {
		t.Error("空内容の項目はfalseを想定しています")
	}
}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	if form.MinLength("some-field", 3) {
		t.Error("存在しない項目の検証結果はNGを想定しています")
	}

	postedData := url.Values{}
	postedData.Add("a", "123456")

	form = New(postedData)
	form.MinLength("a", 10)
	if form.Valid() {
		t.Error("項目が最小桁数より小さいため、検証結果はNGを想定しています")
	}
	es := form.Errors.Get("a")
	if es == "" {
		t.Error("検証結果がNGの場合は、エラーメッセージが取得できる想定です。")
	}

	form = New(postedData)
	form.MinLength("a", 5)
	if !form.Valid() {
		t.Error("OKの検証結果を想定していますが、NGになっています")
	}
	es = form.Errors.Get("a")
	if es != "" {
		t.Error("検証結果がOKの場合は、エラーメッセージが取得できない想定です。", es)
	}
}

func TestForm_IsEmail(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.IsEmail("some-field")
	if form.Valid() {
		t.Error("存在しない項目の検証結果はNGを想定しています")
	}

	postedData := url.Values{}
	postedData.Add("a", "test")
	form = New(postedData)
	form.IsEmail("a")
	if form.Valid() {
		t.Error("メールアドレスフォーマットではないため、検証結果はNGを想定しています")
	}

	postedData = url.Values{}
	postedData.Add("a", "test@test.com")
	form = New(postedData)
	form.IsEmail("a")
	if !form.Valid() {
		t.Error("正しいメールアドレスフォーマットなのに、検証結果がNGになっています")
	}
}
