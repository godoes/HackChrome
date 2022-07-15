package utils

import (
	"syscall"
	"unsafe"
)

type DataBlob struct {
	cbData uint32
	pbData *byte
}

func NewBlob(d []byte) *DataBlob {
	if len(d) == 0 {
		return &DataBlob{}
	}
	return &DataBlob{
		pbData: &d[0],
		cbData: uint32(len(d)),
	}
}

func (b *DataBlob) ToByteArray() []byte {
	d := make([]byte, b.cbData)
	copy(d, (*[1 << 30]byte)(unsafe.Pointer(b.pbData))[:])
	return d
}

func WinDecrypt(data []byte) ([]byte, error) {
	dllCrypt32 := syscall.NewLazyDLL("Crypt32.dll")
	dllKernel32 := syscall.NewLazyDLL("Kernel32.dll")
	procDecryptData := dllCrypt32.NewProc("CryptUnprotectData")
	procLocalFree := dllKernel32.NewProc("LocalFree")

	var outBlob DataBlob
	r, _, err := procDecryptData.Call(uintptr(unsafe.Pointer(NewBlob(data))), 0, 0, 0, 0, 0, uintptr(unsafe.Pointer(&outBlob)))
	if r == 0 {
		return nil, err
	}
	defer func(procLocalFree *syscall.LazyProc, a ...uintptr) {
		_, _, _ = procLocalFree.Call(a...)
	}(procLocalFree, uintptr(unsafe.Pointer(outBlob.pbData)))
	return outBlob.ToByteArray(), nil
}
