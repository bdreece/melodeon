package view

import "github.com/bdreece/melodeon/pkg/spotify/api"

type Model map[string]any

func (m Model) SetUser(user *api.User) {
}

type Options struct {
	TemplateDirectory string            `json:"template_dir"`
	StaticDirectory   string            `json:"static_dir"`
	StaticPrefix      string            `json:"static_prefix"`
	AppDirectory      string            `json:"app_dir"`
	AppPrefix         string            `json:"app_prefix"`
	Layouts           map[string]string `json:"layouts"`
}

var DefaultOptions = Options{
	TemplateDirectory: "web/templates",
	StaticDirectory:   "web/static",
	StaticPrefix:      "/static",
	AppDirectory:      "web/app/dist",
	AppPrefix:         "/dist",
	Layouts: map[string]string{
		"guest": "_guest.gotmpl",
		"host":  "_host.gotmpl",
	},
}
