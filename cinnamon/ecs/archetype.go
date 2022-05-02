package ecs

import (
	"github.com/c1nnam0nbun/cinnamon/ecs/internal/storage"
	"github.com/viant/xunsafe"
	"golang.org/x/exp/slices"
	"reflect"
)

type column struct {
	component componentDescriptor
	data      storage.BlobVec
}

type archetype struct {
	columns  storage.SparseSet[componentId, column]
	entities []int32
}

type Archetype interface {
	addRow(int32)
	addColumn(componentDescriptor)
	contains(id componentId) bool
	components() []componentDescriptor
	containsOnly([]componentDescriptor) bool
	cloneRow(Archetype, int32)
	getColumns() storage.SparseSet[componentId, column]
	getColumn(componentId) *column
	removeRow(int32)
	len() int
	containsAll(types []reflect.Type) bool
	getEntities() []int32
	hash() int32
}

func newArchetype() Archetype {
	arch := &archetype{
		columns:  storage.SparseSet[componentId, column]{},
		entities: make([]int32, 0, 0),
	}
	return arch
}

func (a *archetype) len() int {
	return len(a.entities)
}

func (a *archetype) getEntities() []int32 {
	return a.entities
}

func (a *archetype) hash() int32 {
	var hash int32 = 0
	for _, desc := range a.components() {
		name := desc.t.Name()
		var nameHash int32 = 0
		nameLen := len(name)
		for _, c := range name {
			nameHash += c*31 ^ int32(nameLen-1)
		}
		hash ^= nameHash
	}
	return hash
}

func (a *archetype) contains(id componentId) bool {
	return a.columns.Contains(id)
}

func (a *archetype) containsAll(types []reflect.Type) bool {
	for _, t := range types {
		found := false
		for _, desc := range a.components() {
			if t == desc.t {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func (a *archetype) containsOnly(ids []componentDescriptor) bool {
	components := a.components()
	if len(ids) != len(components) {
		return false
	}
	for _, id := range ids {
		found := false
		for _, desc := range components {
			if desc == id {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func (a *archetype) components() []componentDescriptor {
	components := make([]componentDescriptor, a.columns.GetLen(), a.columns.GetLen())
	for i, col := range a.columns.Dense {
		components[i] = col.component
	}
	return components
}

func (a *archetype) getColumns() storage.SparseSet[componentId, column] {
	return a.columns
}

func (a *archetype) getColumn(id componentId) *column {
	return a.columns.Get(id)
}

func (a *archetype) addRow(id int32) {
	a.entities = append(a.entities, id)
	values := &a.columns.Dense
	for i := 0; i < len(*values); i++ {
		(*values)[i].data.Allocate()
	}
}

func (a *archetype) removeRow(id int32) {
	idx := slices.Index(a.entities, id)
	a.entities = slices.Delete(a.entities, idx, idx+1)
	for _, column := range a.columns.Dense {
		column.data.SwapRemove(int(id))
	}
}

func (a *archetype) cloneRow(other Archetype, id int32) {
	a.addRow(id)
	oldId := slices.Index(other.getEntities(), id)
	newId := slices.Index(a.entities, id)
	for _, column := range other.getColumns().Dense {
		if column.component.size == 0 {
			continue
		}
		xunsafe.Copy(
			a.getColumn(column.component.id).data.Get(newId),
			column.data.Get(oldId),
			column.data.Size,
		)
	}
}

func (a *archetype) addColumn(desc componentDescriptor) {
	a.columns.Set(desc.id, column{
		component: desc,
		data: storage.BlobVec{
			Data: []byte{},
			Size: desc.size,
		},
	})
}
