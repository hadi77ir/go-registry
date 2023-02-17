package registry

type Registry[T any] struct {
	_initialized bool
	items        map[string]T
}

func (r *Registry[T]) init() {
	r.items = make(map[string]T)
	r._initialized = true
}

func (r *Registry[T]) Register(name string, value T) {
	if !r._initialized {
		r.init()
	}
	r.items[name] = value
}
func (r *Registry[T]) Get(name string) (T, bool) {
	if !r._initialized {
		r.init()
	}
	if value, ok := r.items[name]; ok {
		return value, true
	}
	var none T
	return none, false
}

func (r *Registry[T]) Keys() []string {
	if !r._initialized {
		r.init()
	}
	i := 0
	keys := make([]string, len(r.items))
	for k := range r.items {
		keys[i] = k
		i++
	}
	return keys
}
