package ecs

import (
	"github.com/c1nnam0nbun/cinnamon/util"
	"github.com/viant/xunsafe"
	"reflect"
	"unsafe"
)

type query struct {
	world *World
	types []reflect.Type
}

type queryResult map[reflect.Type]any

type QueryResult interface {
	get(reflect.Type) (any, bool)
	set(reflect.Type, any)
}

func (q *queryResult) get(key reflect.Type) (any, bool) {
	r, ok := (*q)[key]
	return r, ok
}

func (q *queryResult) set(key reflect.Type, value any) {
	(*q)[key] = value
}

func MatchQueryResult[T any](result QueryResult) T {
	t := util.TypeOf[T]()
	if v, ok := result.get(t); !ok {
		panic("")
	} else {
		if ptr, ok := v.(unsafe.Pointer); ok {
			return *(*T)(ptr)
		}
		if v, ok := v.(T); ok {
			return v
		}
	}
	panic("")
}

func MatchQueryResultPtr[T any](result QueryResult) *T {
	t := util.TypeOf[T]()
	if v, ok := result.get(t); !ok {
		panic("")
	} else {
		if ptr, ok := v.(unsafe.Pointer); ok {
			return (*T)(ptr)
		}
		if v, ok := v.(T); ok {
			return &v
		}
	}
	panic("")
}

func (q *query) ForEach(fn func(components QueryResult)) {
	result := q.execute()
	for _, c := range result {
		fn(&c)
	}
}

func (q *query) execute() []queryResult {
	result := make([]queryResult, 0, 10)
	idx := 0
	for _, arch := range q.world.archetypes {
		if arch.containsAll(q.types) {
			for _, t := range q.types {
				id := q.world.components.initComponent(t)
				col := arch.getColumn(id)
				for i, e := range arch.getEntities() {
					if len(result) <= i+idx {
						result = append(result, make(queryResult, 0))
						result[i+idx][util.TypeOf[Entity]()] = Entity{
							id:          e,
							archetypeId: arch.getID(),
							world:       q.world,
						}
					}
					if col.component.size != 0 {
						result[i+idx][t] = col.data.Get(i)
					} else {
						result[i+idx][t] = unsafe.Pointer(xunsafe.NewStruct(t))
					}
				}
			}
			idx = len(result)
		}
	}
	return result
}

type Query interface {
	ForEach(func(components QueryResult))
}

func newQuery(world *World, components ...reflect.Type) Query {
	return &query{types: components, world: world}
}
