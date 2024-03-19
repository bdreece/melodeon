package renderer

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/Masterminds/sprig/v3"
	"github.com/labstack/echo/v4"

	"github.com/bdreece/melodeon/pkg/session"
	"github.com/bdreece/melodeon/pkg/view"
)

const (
	userKey  string = "User"
	nonceKey string = "Nonce"
)

var (
	ErrParseTemplate = errors.New("failed to parse templates")
)

type Renderer struct {
	tmpl     *template.Template
	sessions *session.Store
	fs       fs.FS
	layouts  map[string]string
	log      *slog.Logger
}

// Render implements echo.Renderer.
func (r *Renderer) Render(w io.Writer, name string, data any, c echo.Context) error {
	t, err := r.tmpl.Clone()
	if err != nil {
		return fmt.Errorf("failed to clone templates: %w", err)
	}

	templates := r.evaulateTemplates(name)
	t, err = t.ParseFS(r.fs, templates...)
	if err != nil {
		return errors.Join(ErrParseTemplate, err)
	}

	model, ok := data.(view.Model)
	if !ok {
		return t.ExecuteTemplate(w, name, data)
	}

    sess, err := r.sessions.Get(c, session.DefaultCookie)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError)
    }

    if user := sess.User(); user != nil {
        model[userKey] = user
    }

    nonceSess, err := r.sessions.Get(c, session.NonceCookie)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError)
    }
    if nonce := nonceSess.Nonce(); nonce != nil {
        model[nonceKey] = *nonce
    }

	return t.ExecuteTemplate(w, name, model)
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

func New(store *session.Store, log *slog.Logger, opts *view.Options) (*Renderer, error) {
	fs := os.DirFS(opts.TemplateDirectory)
	tmpl, err := template.New("").
		Funcs(sprig.FuncMap()).
		ParseFS(fs, "*.gotmpl")
	if err != nil {
		return nil, errors.Join(ErrParseTemplate, err)
	}

	return &Renderer{tmpl, store, fs, opts.Layouts, log}, nil
}

var _ echo.Renderer = (*Renderer)(nil)
