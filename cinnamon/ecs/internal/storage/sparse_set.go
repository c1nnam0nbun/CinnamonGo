package storage

import (
	"github.com/c1nnam0nbun/cinnamon/util"
	"golang.org/x/exp/slices"
	"reflect"
)

type Index interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type SparseSet[I Index, V any] struct {
	Sparse  []I
	Dense   []V
	Indices []I
	Len     int
}

func (s *SparseSet[I, V]) GetLen() int {
	return s.Len
}

func (s *SparseSet[I, V]) Set(index I, value V) {
	if int(index) >= len(s.Sparse) {
		s.Sparse = Resize(s.Sparse, int(index)+1, -1)
		s.add(index, value)
		return
	}
	if s.Sparse[index] == -1 {
		s.add(index, value)
		return
	}
	s.Dense[int(s.Sparse[index])] = value
}

func (s *SparseSet[I, V]) add(index I, value V) {
	s.Sparse[index] = I(len(s.Indices))
	s.Indices = append(s.Indices, index)
	s.Dense = append(s.Dense, value)
	s.Len++
}

func (s *SparseSet[I, V]) Get(index I) *V {
	if int(index) >= len(s.Sparse) {
		return nil
	}
	idx := s.Sparse[index]
	if idx == -1 {
		return nil
	}
	return &s.Dense[idx]
}

func (s *SparseSet[I, V]) Remove(index I) {
	if len(s.Sparse) <= int(index) {
		return
	}
	last := s.Indices[len(s.Indices)-1]
	idx := s.Sparse[index]
	s.Indices[len(s.Indices)-1], s.Indices[s.Sparse[index]] = s.Indices[s.Sparse[index]], s.Indices[len(s.Indices)-1]
	s.Sparse[last], s.Sparse[index] = s.Sparse[index], s.Sparse[last]
	s.Dense[int(idx)], s.Dense[len(s.Indices)-1] = s.Dense[len(s.Indices)-1], s.Dense[int(idx)]
	s.Indices = slices.Delete(s.Indices, len(s.Indices)-1, len(s.Indices))
	s.Dense = slices.Delete(s.Dense, len(s.Dense)-1, len(s.Dense))
	s.Sparse[index] = -1
	s.Len--
}

func (s *SparseSet[I, V]) Iterate() []V {
	return s.Dense
}

func (s *SparseSet[I, V]) Contains(index I) bool {
	if int(index) >= len(s.Sparse) {
		return false
	}
	return s.Sparse[index] != -1
}

func Resize[I any](slice []I, min int, def I) []I {
	ln := util.Max(len(slice)*2, min)
	newSlice := reflect.MakeSlice(reflect.TypeOf(slice), ln, ln).Interface().([]I)
	copy(newSlice, slice)
	for i := len(slice); i < len(newSlice); i++ {
		newSlice[i] = def
	}
	return newSlice
}
