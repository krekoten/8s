package asm

type OpCode byte

const (
	OpCodeNoOp OpCode = iota
	OpCodeMoveRight
	OpCodeMoveLeft
	OpcodeIncreaseBy
	OpCodeDecreaseBy
	OpCodeInput
	OpCodeOutput
	OpCodeJmpZ
	OpCodeJmpNZ
)
