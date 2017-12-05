package registry

import "github.com/broci/chronicles/ui/component"

type Registry struct {
	components map[string]component.Component
}

func New() *Registry {
	return &Registry{
		components: make(map[string]component.Component),
	}
}

func (r *Registry) Register(name string, c component.Component) {
	r.components[name] = c
}
