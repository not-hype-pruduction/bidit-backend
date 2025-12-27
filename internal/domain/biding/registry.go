package biding

type Registry struct {
	systems map[string]BiddingSystem
}

func NewRegistry(systems ...BiddingSystem) *Registry {
	r := &Registry{systems: make(map[string]BiddingSystem)}
	for _, s := range systems {
		r.systems[s.Name()] = s
	}
	return r
}

func (r *Registry) Get(name string) (BiddingSystem, error) {
	s, ok := r.systems[name]
	if !ok {
		return nil, ErrInvalidSystemName
	}
	return s, nil
}
