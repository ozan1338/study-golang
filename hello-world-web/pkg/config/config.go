package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// AppConfig hold the application config
type AppConfig struct {
	TemplateChache map[string]*template.Template
	UseCache	bool
	InfoLog *log.Logger
	InProdution bool
	Session *scs.SessionManager
}