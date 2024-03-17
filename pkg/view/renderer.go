package view

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"os"

	"github.com/Masterminds/sprig/v3"
	"github.com/labstack/echo/v4"

    "github.com/bdreece/melodeon/pkg/session"
)

var ErrParseTemplate = errors.New("failed to parse templates")

type Renderer struct {
	tmpl *template.Template
    store *session.Store
	fs   fs.FS
}

// Render implements echo.Renderer.
func (r *Renderer) Render(w io.Writer, name string, data any, c echo.Context) error {
	t, err := r.tmpl.Clone()
	if err != nil {
		return fmt.Errorf("failed to clone templates: %w", err)
	}

	t, err = t.ParseFS(r.fs, name)
	if err != nil {
		return errors.Join(ErrParseTemplate, err)
	}

	return t.ExecuteTemplate(w, name, data)
}

func NewRenderer(store *session.Store, opts *Options) (*Renderer, error) {
	fs := os.DirFS(opts.TemplateDirectory)
	tmpl, err := template.New("").
		Funcs(sprig.FuncMap()).
		ParseFS(fs, "*.gotmpl")
	if err != nil {
		return nil, errors.Join(ErrParseTemplate, err)
	}

	return &Renderer{tmpl, store, fs}, nil
}

var _ echo.Renderer = (*Renderer)(nil)
