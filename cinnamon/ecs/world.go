package ecs

import (
	"github.com/c1nnam0nbun/cinnamon/ecs/internal/storage"
	"math"
	"reflect"
)

type World struct {
	archetypes map[int32]Archetype
	components components
	resources  resources
	Scheduler
}

func NewWorld() *World {
	archetypes := make(map[int32]Archetype, 1)
	archetypes[math.MaxInt32] = newArchetype()
	w := &World{
		archetypes: archetypes,
		components: components{
			components:      make([]componentId, 0),
			indices:         make(map[reflect.Type]componentId, 0),
			resourceIndices: make(map[reflect.Type]componentId, 0),
		},
		resources: resources{
			data: storage.ComponentSparseSet[componentId, any]{},
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
	w.archetypes[arch.hash()] = arch
	return arch
}

func (w *World) InitResource(res Resource) {
	w.resources.initResource(componentDescriptor{
		id:   w.components.initResource(res.t),
		t:    res.t,
		size: int(res.t.Size()),
	}, res.data)
}

func (w *World) Query(components ...reflect.Type) Query {
	return newQuery(w, components...)
}
