package main // import "using-library"

import (
	"syscall"
)

func main() {
	fn := syscall.NewLazyDLL("hello.so").NewProc("Hello").Call

	fn()
}
