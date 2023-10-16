package main

type trackCustomData struct {
	trackAudioConfig
	trackVideoConfig
}

type trackAudioConfig struct {
	samplingFrequency float64
	channelConfig     uint16
	bitDepth          uint16
}

type trackVideoConfig struct {
	videoWidth  uint16
	videoHeight uint16
}
