package forms

type errors map[string][]string

// Add 指定された項目のエラーメッセージを追加
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get 指定された項目の一番目エラーメッセージを取得
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}
