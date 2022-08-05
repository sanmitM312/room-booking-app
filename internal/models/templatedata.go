package models

import "github.com/sanmitM312/room-booking-app/internal/forms"

// TemplateData holds data sent from handlers to templates
type TemplateData struct{
	StringMap map[string]string
	IntMap 	  map[string]int
	FloatMap  map[string]float32
	Data 	  map[string]interface{}
	CSRFToken string
	Flash	  string // flash message to user
	Warning   string
	Error	  string
	Form 	  *forms.Form
}