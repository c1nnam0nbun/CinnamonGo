package storage

import (
	"github.com/c1nnam0nbun/cinnamon/util"
	"github.com/viant/xunsafe"
	"golang.org/x/exp/slices"
	"unsafe"
)

type BlobVec struct {
	Data []byte
	Size int
	len  int
}

func (bv *BlobVec) Add(valuePtr unsafe.Pointer) {
	for i := 0; i < bv.Size; i++ {
		b := xunsafe.AsUint8(unsafe.Add(valuePtr, i))
		bv.Data = append(bv.Data, b)
	}
	bv.len++
}

func (bv *BlobVec) Allocate() []byte {
	bv.Data = Resize(bv.Data, util.Max(bv.len*bv.Size, bv.Size), 0)
	bv.len++
	return bv.Data
}

func (bv *BlobVec) Get(index int) unsafe.Pointer {
	return unsafe.Pointer(&bv.Data[index*bv.Size])
}

func (bv *BlobVec) SwapRemove(index int) {
	idxLast := util.Max(len(bv.Data)-bv.Size-1, 0)
	if idxLast != 0 {
		xunsafe.Copy(
			unsafe.Pointer(&bv.Data[index*bv.Size]),
			unsafe.Pointer(&bv.Data[idxLast]),
			bv.Size,
		)
	}
	bv.Data = slices.Delete(bv.Data, idxLast, idxLast+bv.Size)
	bv.len--
}

func (bv *BlobVec) Set(index int, valuePtr unsafe.Pointer) {
	xunsafe.Copy(
		unsafe.Pointer(&bv.Data[index*bv.Size]),
		valuePtr,
		bv.Size,
	)
}
