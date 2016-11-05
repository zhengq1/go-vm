package vm

type VMState byte

const (
	NONE = 0
	HALT = 1 << 0
	FAULT = 1 << 1
)
