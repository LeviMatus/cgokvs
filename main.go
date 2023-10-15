package main

// #cgo CFLAGS: -Isrc/include -Idependency/libkvspic/kvspic-src/src/client/include -Idependency/libkvspic/kvspic-src/src/common/include -Idependency/libkvspic/kvspic-src/src/utils/include -Idependency/libkvspic/kvspic-src/src/mkvgen/include -Idependency/libkvspic/kvspic-src/src/view/include -Idependency/libkvspic/kvspic-src/src/heap/include -Idependency/libkvspic/kvspic-src/src/state/include
// #cgo LDFLAGS: -L${SRCDIR}/build -lcproducer
// #include <stdio.h>
// #include <com/amazonaws/kinesis/video/cproducer/Include.h>
import "C"

func main() {
	C.puts(C.CString("Hello, World!"))
}
