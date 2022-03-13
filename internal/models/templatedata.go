package models

import "github.com/techarm/go-bookings/internal/forms"

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	Float     map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Info      string
	Warning   string
	Error     string
	Form      *forms.Form
}
