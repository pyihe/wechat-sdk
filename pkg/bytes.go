package pkg

import "unsafe"

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func BytesContain(ele byte, b []byte) bool {
	for _, v := range b {
		if v == ele {
			return true
		}
	}
	return false
}
