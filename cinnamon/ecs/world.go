package ecs

import (
	"math"
	"reflect"
)

type World struct {
	archetypes map[int32]Archetype
	components components
	Scheduler
}

func NewWorld() *World {
	archetypes := make(map[int32]Archetype, 1)
	archetypes[math.MaxInt32] = newArchetype()
	w := &World{
		archetypes: archetypes,
		components: components{
			components: make([]componentId, 0),
			indices:    make(map[reflect.Type]componentId, 0),
		},
	}
	w.Scheduler = NewScheduler(w)
	return w
}

func (w *World) CreateEntity() Entity {
	e := newEntity(w)
	w.archetypes[math.MaxInt32].addRow(e.id)
	return e
}

func (w *World) initArchetype(components ...componentDescriptor) Archetype {
	arch := newArchetype()
	for _, c := range components {
		arch.addColumn(c)
	}
	for i := 0; i < 3; i++ {
		println(arch.getID())
	}
	w.archetypes[arch.getID()] = arch
	return arch
}

func (w *World) Query(components ...reflect.Type) Query {
	return newQuery(w, components...)
}
