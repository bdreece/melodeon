package app

import (
	"fmt"
	"os"

	"github.com/bdreece/melodeon/pkg/config"
	"github.com/bdreece/melodeon/pkg/logger"
	"github.com/bdreece/melodeon/pkg/session"
	"github.com/bdreece/melodeon/pkg/spotify"
	"github.com/bdreece/melodeon/pkg/store"
	"github.com/bdreece/melodeon/pkg/view"
	"go.uber.org/dig"
)

type Builder struct{ container *dig.Container }

func (b *Builder) With(constructor any, opts ...dig.ProvideOption) *Builder {
	if err := b.container.Provide(constructor, opts...); err != nil {
		b.quit(err)
	}

	return b
}

func (b *Builder) Build() *App {
	_ = b.container.Provide(New)

	ch := make(chan *App, 1)
	defer close(ch)
	if err := b.container.Invoke(func(app *App) {
		ch <- app
	}); err != nil {
		panic(err)
	}

	return <-ch
}

func (b *Builder) quit(err error) {
	fmt.Fprintf(os.Stderr, "unexpected error during startup: %v", err)
	_ = dig.Visualize(b.container, os.Stdout, dig.VisualizeError(err))
	os.Exit(1)
}

func NewBuilder(cfgpath string) *Builder {
	type options struct {
		dig.Out

		App     *config.AppOptions
		Logger  *logger.Options
		Session *session.Options
		Spotify *spotify.Options
		Store   *store.Options
		View    *view.Options
	}

	builder := Builder{dig.New(dig.RecoverFromPanics())}
	return builder.With(func() (opts options, err error) {
		cfg, err := config.Load(cfgpath)
		if err != nil {
			return
		}

		opts.App = &cfg.AppOptions
		opts.Logger = &cfg.Logger
		opts.Session = &cfg.Session
		opts.Spotify = &cfg.Spotify
		opts.Store = &cfg.Store
		opts.View = &cfg.View

		return
	})
}
