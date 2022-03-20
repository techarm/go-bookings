package forms

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"net/url"
	"strings"
)

// Form 画面Formデータと検証エラー情報
type Form struct {
	url.Values
	Errors errors
}

// New 画面Form初期化処理
func New(data url.Values) *Form {
	return &Form{
		data,
		map[string][]string{},
	}
}

// Valid Form検証OKの場合はtrueを返す、それ以外の場合はfalseを返す
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// Required 指定された複数項目の必須チェックを行う
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "値を入力してください。")
		}
	}
}

// Has 指定された項目が空かどうかをチェックする
func (f *Form) Has(field string) bool {
	x := f.Get(field)
	if x == "" {
		return false
	}
	return true
}

// MinLength 指定された桁数の最小桁チェック
func (f *Form) MinLength(field string, length int) bool {
	x := f.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("値を%d桁以上を入力してください。", length))
		return false
	}
	return true
}

// IsEmail メールアドレスのフォーマットチェック
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "正しいメールアドレスを入力してください。")
	}
}
