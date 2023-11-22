package config

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the app config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	//this parameter sets Cookies depends we are running as a production or development server
	InProduction bool
	Session *scs.SessionManager
}
