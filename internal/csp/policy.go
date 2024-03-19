package csp

import "strings"

const (
	self         string = "'self'"
	unsafeInline string = "'unsafe-inline'"
	jsdelivr     string = "cdn.jsdelivr.net"
	googleapis   string = "fonts.googleapis.com"
	gstatic      string = "fonts.gstatic.com"
	scdn         string = "sdk.scdn.co"
	any          string = "*"
)

type Policy struct {
	DefaultSrc []string
	ScriptSrc  []string
	StyleSrc   []string
	FontSrc    []string
	FrameSrc   []string
	ImgSrc     []string
}

func (p Policy) String() string {
	b := new(strings.Builder)
	_, _ = b.WriteString("default-src " + strings.Join(p.DefaultSrc, " ") + "; ")
	_, _ = b.WriteString("script-src " + strings.Join(p.ScriptSrc, " ") + "; ")
	_, _ = b.WriteString("style-src " + strings.Join(p.StyleSrc, " ") + "; ")
	_, _ = b.WriteString("frame-src " + strings.Join(p.FrameSrc, " ") + "; ")
	_, _ = b.WriteString("font-src " + strings.Join(p.FontSrc, " ") + "; ")
	_, _ = b.WriteString("img-src " + strings.Join(p.ImgSrc, " "))
	return b.String()
}

var DefaultPolicy = Policy{
	DefaultSrc: []string{self},
	ScriptSrc:  []string{jsdelivr, scdn},
	StyleSrc:   []string{self, unsafeInline, googleapis, jsdelivr},
	FrameSrc:   []string{scdn},
	FontSrc:    []string{gstatic, jsdelivr},
	ImgSrc:     []string{any},
}

