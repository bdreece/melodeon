GO 	   		= go
GOFLAGS 	= -v
AIR 		= $(GO) run github.com/cosmtrek/air@latest
AIRFLAGS 	= -root "." \
			  -build.bin "bin/melodeon" \
			  -build.cmd "make" \
			  -build.exclude_dir "web/app/dist,bin" \
			  -build.include_ext "go,gotmpl,js,scss,json,yml" \
			  -build.kill_delay "0.5s" \
			  -build.send_interrupt "true" \
			  -screen.clear_on_rebuild "true" \
			  -tmp_dir "tmp"

INSTALL 		= install
INSTALL_PROGRAM = $(INSTALL)
INSTALL_DATA 	= $(INSTALL) -m 0644

prefix 		= /usr/local
exec_prefix = $(prefix)
datarootdir = $(prefix)/share
datadir 	= $(datarootdir)
bindir 		= $(exec_prefix)/bin
docdir		= $(datadir)/doc/notable

src := $(shell find . -type f -name "*.go") \
	   $(shell find . -type f -name "*.gotmpl")

cmddir 		:= ./cmd/melodeon
targetdir 	:= ./bin
webappdir 	:= ./web/app
assets 		:= $(shell find $(webappdir) -type f -name "*.js") \
			   $(shell find $(webappdir) -type f -name "*.scss")
dist 		:= $(webappdir)/dist/melodeon.umd.cjs $(webappdir)/dist/style.css
