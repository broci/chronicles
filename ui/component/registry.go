package component

type Registry struct {
	components map[string]Component
}

func NewRegistry() *Registry {
	return &Registry{
		components: make(map[string]Component),
	}
}

func (r *Registry) Register(name string, c Component) {
	r.components[name] = c
}

func (r *Registry) Get(name string) Component {
	return r.components[name]
}
