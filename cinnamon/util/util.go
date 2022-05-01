package util

import (
	"reflect"
	"unsafe"
)

func TypeOf[T any]() reflect.Type {
	return reflect.TypeOf((*T)(nil)).Elem()
}

func SizeOf[T any]() uintptr {
	return unsafe.Sizeof(TypeOf[T]())
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
