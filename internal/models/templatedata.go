package models

import "github.com/TartuDen/webPage4_starting_project/internal/forms"

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	//cross site request forgery token (security token)
	CSRFToken string
	//dif kind of messages
	Flash   string
	Warning string
	Error   string
	Form *forms.Form
}
