package ecs

import (
	"fmt"
	"golang.org/x/exp/slices"
	"math"
	"reflect"
	"unsafe"
)

type Entity struct {
	id          int32
	archetypeId int32
	world       *World
}

var entityCounter int32 = 0

func (e Entity) String() string {
	return fmt.Sprintf("{%d}", e.id)
}

func newEntity(world *World) Entity {
	entity := Entity{
		id:          entityCounter,
		archetypeId: math.MaxInt32,
		world:       world,
	}
	entityCounter++
	return entity
}

func AddComponent[T any](e *Entity, component T) error {
	archetype := e.world.archetypes[e.archetypeId]
	t := reflect.TypeOf(component)
	id := e.world.components.initComponent(t)
	size := int(unsafe.Sizeof(component))
	if archetype.contains(id) {
		return fmt.Errorf("entity already contains component '%s'", t.Name())
	}
	components := append(archetype.components(), componentDescriptor{id, t, size})
	var newArch Archetype
	for _, arch := range e.world.archetypes {
		if arch.containsOnly(components) {
			newArch = arch
		}
	}
	if newArch == nil {
		newArch = e.world.initArchetype(components...)
	}
	e.archetypeId = newArch.hash()
	newArch.cloneRow(archetype, e.id)
	if size > 0 {
		newArch.getColumn(id).data.Set(slices.Index(newArch.getEntities(), e.id), unsafe.Pointer(&component))
	}
	archetype.removeRow(e.id)
	if archetype.len() == 0 {
		delete(e.world.archetypes, archetype.hash())
	}
	return nil
}
