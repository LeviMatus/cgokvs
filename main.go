package main

// #cgo CFLAGS: -Isrc/include -Idependency/libkvspic/kvspic-src/src/client/include -Idependency/libkvspic/kvspic-src/src/common/include -Idependency/libkvspic/kvspic-src/src/utils/include -Idependency/libkvspic/kvspic-src/src/mkvgen/include -Idependency/libkvspic/kvspic-src/src/view/include -Idependency/libkvspic/kvspic-src/src/heap/include -Idependency/libkvspic/kvspic-src/src/state/include
// #cgo LDFLAGS: -L${SRCDIR}/build -lcproducer
// #include <stdio.h>
// #include <com/amazonaws/kinesis/video/cproducer/Include.h>
import "C"

type StreamInfo struct {
	// Version of the struct
	version uint32

	// Stream name - human readable. Null terminated.
	// Should be unique per AWS account.
	name [129]byte

	// Number of tags associated with the stream
	tagCount uint32

	// Stream tags array
	tags []Tag

	// Stream retention period in 100ns - 0 for no retention. Retention should be greater than 1 hour as
	// the service-side accepts the retention policy in an hour units.
	retention uint64

	// KMS key id ARN
	kmsKeyId [2049]byte

	// Stream capabilities
	streamCaps []StreamCap
}

func main() {
	C.puts(C.CString("Hello, World!"))
	C.createRealtimeVideoStreamInfoProvider(C.CString("levis-test"), 100, 100, nil)
}
