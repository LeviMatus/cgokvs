package main

type Tag struct {
	// Version of the struct
	version uint32

	// Tag name - null terminated
	name []string // pointer to a string with MAX_TAG_NAME_LEN chars max including the NULL terminator

	// Tag value - null terminated
	value []string // pointer to a string with MAX_TAG_VALUE_LEN chars max including the NULL terminator
}

type StreamingType int

const (
	// Realtime mode for minimal latency
	STREAMING_TYPE_REALTIME StreamingType = iota
	// Near-realtime mode for predefined latency
	STREAMING_TYPE_NEAR_REALTIME
	// Offline upload mode
	STREAMING_TYPE_OFFLINE
)

type StreamCap struct {
	// Streaming type
	streamingType StreamingType

	// Stream content type - null terminated.
	contentType [128 + 1]byte

	// Whether the bitrate can change in mid-stream.
	adaptive bool

	// Max latency tolerance in time units. Can be STREAM_LATENCY_PRESSURE_CHECK_SENTINEL for realtime streaming.
	maxLatency uint64

	// Duration of the fragment/cluster. Can be FRAGMENT_KEY_FRAME_DURATION_SENTINEL if based on IDR/key frames.
	fragmentDuration uint64

	// Whether to create fragments on the IDR/key frame boundary or based on duration.
	keyFrameFragmentation bool

	// Whether to use frame timecodes.
	frameTimecodes bool

	// Whether the clusters will have absolute or relative timecodes
	absoluteFragmentTimes bool

	// Whether the application ACKs are required
	fragmentAcks bool

	// Whether to recover after an error occurred
	recoverOnError bool

	// Specify the NALs adaptation flags as defined in NAL_ADAPTATION_FLAGS enumeration
	// The adaptation will be applied to all tracks
	nalAdaptationFlags uint32

	// Average stream bandwidth requirement in bits per second
	avgBandwidthBps uint32

	// Number of frames per second. Will use the defaults if 0.
	frameRate uint32

	// Duration of content to keep in 100ns in storage before purging.
	// 0 for default values to be calculated based on replay buffer, etc..
	bufferDuration uint64

	// Duration of content in 100ns to re-transmit after reconnection.
	// 0 if the latest frame is to be re-transmitted in Realtime mode
	// For Near-Realtime mode or offline mode it will be ignored.
	// If we receive non "dead host" error in connection termination event
	// and the ACKs are enabled then we can rollback to an earlier timestamp
	// as the host has already received the fragment.
	replayDuration uint64

	// Duration to check back for the connection staleness.
	// Can be CONNECTION_STALENESS_DETECTION_SENTINEL to skip the check.
	// If we haven't received any buffering ACKs and the delta between the current
	// and the last buffering ACK is greater than this duration the customer
	// provided optional callback will be executed.
	connectionStalenessDuration uint64

	// Timecode scale to use generating the packaging.
	// NOTE: Specifying DEFAULT_TIMECODE_SCALE_SENTINEL will imply using
	// default timecode for the packaging.
	timecodeScale uint64

	// Whether to recalculate metrics at runtime with slight increasing performance hit.
	recalculateMetrics bool

	// Segment UUID. If specified it should be MKV_SEGMENT_UUID_LEN long. Specifying NULL will generate random UUID
	segmentUuid []byte

	// Array of TrackInfo containing track metadata
	trackInfoList []TrackInfo

	// Number of TrackInfo in trackInfoList
	trackInfoCount uint32

	// ------------------------------- V0 compat ----------------------

	// How incoming frames are reordered
	frameOrderingMode FrameOrderMode

	// ------------------------------- V1 compat ----------------------

	// Content store pressure handling policy
	storePressurePolicy ContentStorePressurePolicy

	// Content view overflow handling policy
	viewOverflowPolicy ContentViewOverflowPolicy

	// ------------------------------ V2 compat -----------------------
	// Enable / Disable stream creation if describe call fails
	allowStreamCreation bool
}

type MkvTrackInfoType byte

const (
	MKV_TRACK_INFO_TYPE_VIDEO  MkvTrackInfoType = 0x01
	MKV_TRACK_INFO_TYPE_AUDIO  MkvTrackInfoType = 0x02
	MKV_TRACK_INFO_TYPE_UNKOWN MkvTrackInfoType = 0x03
)

type TrackInfo struct {
	version uint32

	// Unique Identifier for TrackInfo
	trackId uint64

	// Codec ID of the stream. Null terminated.
	codecId [32 + 1]byte

	// Human readable track name. Null terminated.
	trackName [32 + 1]byte

	// Size of the codec private data in bytes. Can be 0 if no CPD is used.
	codecPrivateDataSize uint32

	// Codec private data. Can be NULL if no CPD is used. Allocated in heap.
	codecPrivateData []byte

	// Track's content type.
	trackType MkvTrackInfoType

	// Track type specific data.
	trackCustomData trackCustomData
}
