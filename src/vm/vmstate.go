package vm

type VMState byte

const (
	NONE VMState = 0
	HALT VMState = 1 << 0
	FAULT VMState = 1 << 1
)
