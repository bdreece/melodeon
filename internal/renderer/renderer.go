package renderer

import (
	"html/template"
	"io"
	"io/fs"
	"os"

	"github.com/Masterminds/sprig/v3"
	"github.com/labstack/echo/v4"
)

type Renderer struct {
	tmpl *template.Template
	fs   fs.FS
}

func (r Renderer) Render(w io.Writer, name string, data any, c echo.Context) error {
	t, err := template.Must(r.tmpl.Clone()).
		Funcs(funcMap(c)).
		ParseFS(r.fs, name)

	if err != nil {
		return err
	}

	return t.ExecuteTemplate(w, name, data)
}

func New(cfg *Config) (*Renderer, error) {
	fs := os.DirFS(cfg.TemplateDirectory)
	tmpl, err := template.New("").
		Funcs(sprig.FuncMap()).
		Funcs(funcMap(nil)).
		ParseFS(fs, "**/*.gotmpl")

	if err != nil {
		return nil, err
	}

	return &Renderer{tmpl, fs}, nil
}
