package view

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log/slog"
	"os"
	"strings"

	"github.com/Masterminds/sprig/v3"
	"github.com/labstack/echo/v4"

    "github.com/bdreece/melodeon/pkg/session"
)

var (
	ErrParseTemplate = errors.New("failed to parse templates")
)

type Renderer struct {
	tmpl    *template.Template
	store   *session.Store
	fs      fs.FS
	layouts map[string]string
	log     *slog.Logger
}

// Render implements echo.Renderer.
func (r *Renderer) Render(w io.Writer, name string, data any, c echo.Context) error {
	t, err := r.tmpl.Clone()
	if err != nil {
		return fmt.Errorf("failed to clone templates: %w", err)
	}

    templates := r.evaulateTemplates(name)
	t, err = t.Funcs(funcMap(c)).ParseFS(r.fs, templates...)
	if err != nil {
		return errors.Join(ErrParseTemplate, err)
	}

	if v, ok := data.(Model); ok {
		session, err := r.store.Get(c, session.Name)
		if err != nil {
			return fmt.Errorf("failed to get session: %w", err)
		}

		user := session.User()
		if user != nil {
			r.log.Debug("injecting user model",
				slog.String("display_name", user.DisplayName))

			v["User"] = user
			return t.ExecuteTemplate(w, name, v)
		}
	}

	return t.ExecuteTemplate(w, name, data)
}

func (r Renderer) evaulateTemplates(name string) []string {
    templates := []string{}
    for prefix, layout := range r.layouts {
        if strings.HasPrefix(name, prefix) {
            templates = append(templates, layout)
        }
    }

    return append(templates, name)
}

func NewRenderer(store *session.Store, log *slog.Logger, opts *Options) (*Renderer, error) {
	fs := os.DirFS(opts.TemplateDirectory)
	tmpl, err := template.New("").
		Funcs(sprig.FuncMap()).
        Funcs(funcMap(nil)).
		ParseFS(fs, "*.gotmpl")
	if err != nil {
		return nil, errors.Join(ErrParseTemplate, err)
	}

	return &Renderer{tmpl, store, fs, opts.Layouts, log}, nil
}

var _ echo.Renderer = (*Renderer)(nil)
