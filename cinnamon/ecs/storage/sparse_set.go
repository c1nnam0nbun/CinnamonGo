package storage

import "github.com/c1nnam0nbun/cinnamon/ecs/internal/storage"

type SparseSet[I storage.Index, V any] interface {
	Set(index I, value V)
	Get(index I) *V
	Remove(index I)
	Contains(index I) bool
	Iterate() []V
	GetLen() int
}

func NewSparseSet[I storage.Index, V any]() SparseSet[I, V] {
	return &storage.SparseSet[I, V]{}
}
