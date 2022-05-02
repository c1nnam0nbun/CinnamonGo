package storage

import (
	"golang.org/x/exp/slices"
	"unsafe"
)

type ComponentSparseSet[I Index, V any] struct {
	Sparse  []I
	Dense   BlobVec
	Indices []I
	Len     int
}

func (s *ComponentSparseSet[I, V]) GetLen() int {
	return s.Len
}

func (s *ComponentSparseSet[I, V]) Set(index I, value V) {
	if int(index) >= len(s.Sparse) {
		s.Sparse = Resize(s.Sparse, int(index)+1, -1)
		s.add(index, value)
		return
	}
	if s.Sparse[index] == -1 {
		s.add(index, value)
		return
	}
	s.Dense.Set(int(s.Sparse[index]), unsafe.Pointer(&value))
}

func (s *ComponentSparseSet[I, V]) add(index I, value V) {
	s.Sparse[index] = I(len(s.Indices))
	s.Indices = append(s.Indices, index)
	if value == nil {

	}
	s.Dense.Add(unsafe.Pointer(&value))
	s.Len++
}

func (s *ComponentSparseSet[I, V]) Get(index I) *V {
	if int(index) >= len(s.Sparse) {
		return nil
	}
	idx := s.Sparse[index]
	if idx == -1 {
		return nil
	}
	return (*V)(s.Dense.Get(int(idx)))
}

func (s *ComponentSparseSet[I, V]) Remove(index I) {
	if len(s.Sparse) <= int(index) {
		return
	}
	last := s.Indices[len(s.Indices)-1]
	idx := s.Sparse[index]
	s.Indices[len(s.Indices)-1], s.Indices[s.Sparse[index]] = s.Indices[s.Sparse[index]], s.Indices[len(s.Indices)-1]
	s.Sparse[last], s.Sparse[index] = s.Sparse[index], s.Sparse[last]
	s.Dense.SwapRemove(int(idx))
	s.Indices = slices.Delete(s.Indices, len(s.Indices)-1, len(s.Indices))
	s.Sparse[index] = -1
	s.Len--
}

//func (s *ComponentSparseSet[I, V]) Iterate() []V {
//	return s.Dense
//}

func (s *ComponentSparseSet[I, V]) Contains(index I) bool {
	if int(index) >= len(s.Sparse) {
		return false
	}
	return s.Sparse[index] != -1
}
