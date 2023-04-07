//go:build !windows

package tiktokengo

//go:generate bash -c "cd tiktokenrs && cargo build --release"

/*
#cgo linux LDFLAGS: ${SRCDIR}/tiktokenrs/target/release/libtiktoken.a -ldl
#cgo darwin LDFLAGS: ${SRCDIR}/tiktokenrs/target/release/libtiktoken.a -framework Security -framework CoreFoundation

#include <stdlib.h>
extern unsigned int get_qtd_tokens(const char*);
*/
import "C"
import (
	"unsafe"
)

// bpe is only cl100k_base
func CountTokens(prompt string) int {
	p := C.CString(prompt)
	count := C.get_qtd_tokens(p)
	C.free(unsafe.Pointer(p))
	return int(count)
}
