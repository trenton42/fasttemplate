//go:build !appengine

package fasttemplate

import (
	"unsafe"
)

func unsafeBytes2String(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}

func unsafeString2Bytes(s string) (b []byte) {
	if len(s) == 0 {
		return []byte{}
	}
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
