package ecs

import (
	"github.com/c1nnam0nbun/cinnamon/ecs/internal/storage"
	"github.com/c1nnam0nbun/cinnamon/util"
	"reflect"
)

type resource struct {
	desc componentDescriptor
	data storage.BlobVec
}

type Resource struct {
	data any
	t    reflect.Type
}

type Resources []Resource

type resources struct {
	data storage.ComponentSparseSet[componentId, any]
}

func NewResource[T any](res any) Resource {
	return Resource{res, util.TypeOf[T]()}
}

func (r *resources) initResource(desc componentDescriptor, data any) {
	if data == nil {
	}
	r.data.Set(desc.id, data)
}
