# Shared libarary test on Windows

## Not support below for Windows
```sh
go build -buildmode=plugin
go build -buildmode=shared
go build -linkshared
```
* Also Cgo is disabled when cross compile mode.

## Create shared library
```sh
cd so
go build -buildmode=c-shared -o hello.so
```

## Compile main.go with so
* <b>Not work</b>
* Copy `hello.h`, `hello.so` to `main`
```sh
cd main
go build
```

## Compile using syscall with so
* <b>Not work</b>
* https://github.com/golang/go/issues/22192
* Copy `hello.so` to `main-syscall`
```sh
cd main-syscall
go build
```

## Compile from C with header
* <b>Work</b>
* Copy `hello.h`, `hello.so` to `call-from-c`
```sh
cd call-from-c
gcc main.c
```

## Compile from C using syscall
* <b>Work</b>
* https://github.com/golang/go/issues/22192#issuecomment-335336513
* Copy hello.so to `call-from-c-syscall`
```sh
cd call-from-c-syscall
gcc main.c ./hello.so
```
