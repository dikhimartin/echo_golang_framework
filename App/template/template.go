package template

import (
	"github.com/labstack/echo"
	"html/template"
	"io"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Rendering() {
	renderer := &TemplateRenderer{
		  templates: template.Must(template.ParseFiles(
			"template/header.html",
			"template/layout.html",		
			"template/footer.html",		
		)),
	}
}