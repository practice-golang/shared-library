package main // import "shared-library"

// #cgo CFLAGS: -g -Wall
// #include "hello.h"
import "C"

func main() {
	C.Hello()
}
