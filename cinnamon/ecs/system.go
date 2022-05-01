package ecs

type systemFn func(*World)

type system struct {
	system systemFn
	label  string
	before []string
	after  []string
}

type System interface {
	Label(label string) System
	Before(label string) System
	After(label string) System
	call(world *World)
}

func NewSystem(fn systemFn) System {
	return &system{
		system: fn,
	}
}

func (s *system) call(world *World) {
	s.system(world)
}

func (s *system) Label(label string) System {
	s.label = label
	return s
}

func (s *system) Before(label string) System {
	s.before = append(s.before, label)
	return s
}

func (s *system) After(label string) System {
	s.after = append(s.after, label)
	return s
}
