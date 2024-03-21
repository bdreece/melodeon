package csp

import "strings"

const (
	self         string = "'self'"
	unsafeInline string = "'unsafe-inline'"
	jsdelivr     string = "cdn.jsdelivr.net"
	googleapis   string = "fonts.googleapis.com"
	gstatic      string = "fonts.gstatic.com"
	scdn         string = "sdk.scdn.co"
	dealer       string = "wss://gue1-dealer.spotify.com"
	any          string = "*"
)

type Policy struct {
	DefaultSrc []string
	ConnectSrc []string
	ScriptSrc  []string
	StyleSrc   []string
	FontSrc    []string
	FrameSrc   []string
	ImgSrc     []string
}

func (p Policy) String() string {
	b := new(strings.Builder)
	_, _ = b.WriteString("default-src " + strings.Join(p.DefaultSrc, " ") + "; ")
	_, _ = b.WriteString("connect-src " + strings.Join(p.ConnectSrc, " ") + "; ")
	_, _ = b.WriteString("script-src " + strings.Join(p.ScriptSrc, " ") + "; ")
	_, _ = b.WriteString("style-src " + strings.Join(p.StyleSrc, " ") + "; ")
	_, _ = b.WriteString("frame-src " + strings.Join(p.FrameSrc, " ") + "; ")
	_, _ = b.WriteString("font-src " + strings.Join(p.FontSrc, " ") + "; ")
	_, _ = b.WriteString("img-src " + strings.Join(p.ImgSrc, " "))
	return b.String()
}

var DefaultPolicy = Policy{
	DefaultSrc: []string{self},
	ConnectSrc: []string{self, dealer},
	ScriptSrc:  []string{self, jsdelivr, scdn},
	StyleSrc:   []string{self, unsafeInline, googleapis, jsdelivr},
	FrameSrc:   []string{scdn},
	FontSrc:    []string{gstatic, jsdelivr},
	ImgSrc:     []string{any},
}
