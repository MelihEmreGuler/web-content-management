package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
)

// AppConfig configurations for the application
type AppConfig struct {
	UseCache      bool                          // set to true to use template cache
	TemplateCache map[string]*template.Template // map of template
	InProduction  bool                          // set to true in production
	Session       *scs.SessionManager           // session manager
}
