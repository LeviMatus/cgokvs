package main

type ContentStorePressurePolicy int

const (
	ContentStorePressurePolicyOOM ContentStorePressurePolicy = iota
	ContentStorePressurePolicyDropTailItem
)

type ContentViewOverflowPolicy int

const (
	ContentViewOverflowPoilcyDropTailViewItem ContentViewOverflowPolicy = iota
	ContentViewOverflowPolicyDropUntilFragmentStart
)
