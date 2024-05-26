GO 	   		= go
GOFLAGS 	= -v
NPM         = npm
AIR 		= $(GO) run github.com/cosmtrek/air@latest
AIRFLAGS 	= -root "." \
			  -build.bin "bin/melodeon" \
			  -build.args_bin "-p 3000 -c ./configs/development.yml" \
			  -build.cmd "make" \
			  -build.exclude_dir "bin,tmp,node_modules,web/app/node_modules,web/app/dist" \
			  -build.include_ext "go,gotmpl,js,scss,json,yml" \
			  -build.kill_delay "0.5s" \
			  -build.send_interrupt "true" \
			  -screen.clear_on_rebuild "true" \
			  -tmp_dir "tmp"

cmddir 		:= ./cmd/melodeon
configdir 	:= ./configs
targetdir 	:= ./bin
webdir 		:= ./web
webappdir 	:= $(webdir)/app

src 	:= $(shell find . -type f -name "*.go") \
		   $(shell find . -type f -name "*.gotmpl")
assets 	:= $(shell find $(webappdir) -type f -name "*.js") \
		   $(shell find $(webappdir) -type f -name "*.scss")
dist 	:= $(webappdir)/dist/melodeon.umd.cjs \
		   $(webappdir)/dist/style.css
