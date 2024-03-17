package route

type Route interface {
	Pattern() string
}

type route struct {
	pattern string
}

func (r route) Pattern() string { return r.pattern }

func New(pattern string) Route {
    return route{pattern}
}
