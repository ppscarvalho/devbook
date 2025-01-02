package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// LoadTemplates carrega os templates
func LoadTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

// RenderTemplate renderiza um template
func RenderTemplate(w http.ResponseWriter, nome string, dados interface{}) {
	templates.ExecuteTemplate(w, nome, dados)
}

// EscaparHTML escapa caracteres especiais para evitar ataques XSS
