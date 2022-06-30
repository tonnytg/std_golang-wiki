package main

// #cgo CFLAGS: -g -Wall
// #include <stdlib.h>
// #include "main.h"
import "C"
import (
	"fmt"
	"unsafe"
	"bufio"
)

func main() {
	name := C.CString("Gopher")
	defer C.free(unsafe.Pointer(name))

	year := C.int(2018)

	ptr := C.malloc(C.sizeof_char * 1024)
	defer C.free(unsafe.Pointer(ptr))

	size := C.greet(name, year, (*C.char)(ptr))

	b := C.GoBytes(ptr, size)
	fmt.Println(string(b))
}