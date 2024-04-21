package fastconv

import (
	"unsafe"
)

// Быстрое преобразование слайса байтов в строку.
// Можно передавать nil.
// Нельзя использовать, если bytes может меняться.
func String(bytes []byte) string {
	return unsafe.String(unsafe.SliceData(bytes), len(bytes))
}

// Быстрое преобразование строки в слайс байтов.
// Нельзя использовать, если bytes будет меняться.
func Bytes(str string) []byte {
	return unsafe.Slice(unsafe.StringData(str), len(str))
}

// Возвращает указатель на информацию о типе интерфейса.
func TypePointer(a any) uintptr {
	type emptyInterface struct {
		typ unsafe.Pointer
		ptr unsafe.Pointer
	}

	iface := (*emptyInterface)(unsafe.Pointer(&a))

	return uintptr(iface.typ)
}
