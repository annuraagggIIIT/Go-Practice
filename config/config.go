package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
	"log"
)

// AppConfig holds the app config
type AppConfig struct {
	HttpSecure    bool
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	Session       *scs.SessionManager
}
