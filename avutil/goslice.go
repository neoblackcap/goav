package avutil

import "unsafe"

func PointerToUint8Slice(p unsafe.Pointer, size int) []uint8 {
	if p == nil {
		return nil
	}
	return (*[1 << 30]uint8)(p)[:size]
}

func PointerToUint16Slice(p unsafe.Pointer, size int) []uint16 {
	if p == nil {
		return nil
	}
	return (*[1 << 30]uint16)(p)[:size/2]
}

func PointerToUint32Slice(p unsafe.Pointer, size int) []uint32 {
	if p == nil {
		return nil
	}
	return (*[1 << 30]uint32)(p)[:size/4]
}
