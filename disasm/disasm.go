package disasm

import (
	"fmt"
	"strings"

	"github.com/krekoten/8s/asm"
)

var opcodeToString map[asm.OpCode]string = map[asm.OpCode]string{
	asm.OpCodeMoveLeft:   "OpCodeMoveLeft",
	asm.OpCodeMoveRight:  "OpCodeMoveRight",
	asm.OpCodeInput:      "OpCodeInput",
	asm.OpCodeOutput:     "OpCodeOutput",
	asm.OpCodeNoOp:       "OpCodeNoOp",
	asm.OpcodeIncreaseBy: "OpcodeIncreaseBy",
	asm.OpCodeDecreaseBy: "OpCodeDecreaseBy",
	asm.OpCodeJmpZ:       "OpCodeJmpZ",
	asm.OpCodeJmpNZ:      "OpCodeJmpNZ",
}

type disasm struct {
	code    []asm.OpCode
	nesting int
}

func New(code []asm.OpCode) *disasm {
	return &disasm{code: code}
}

func (d *disasm) Print() string {
	var result strings.Builder

	row := 1
	for i := 0; i < len(d.code); i++ {
		instr := d.code[i]
		result.WriteString(fmt.Sprintf("%v. ", row))
		row += 1
		switch instr {
		case asm.OpCodeMoveLeft, asm.OpCodeMoveRight, asm.OpCodeInput, asm.OpCodeOutput, asm.OpCodeNoOp:
			result.WriteString(fmt.Sprintf("%s%s\n", strings.Repeat(" ", d.nesting), opcodeToString[instr]))
		case asm.OpcodeIncreaseBy, asm.OpCodeDecreaseBy:
			result.WriteString(fmt.Sprintf("%s%s (%v)\n", strings.Repeat(" ", d.nesting), opcodeToString[instr], d.code[i+1]))
			i += 1
		case asm.OpCodeJmpZ:
			result.WriteString(fmt.Sprintf("%s%s (%v)\n", strings.Repeat(" ", d.nesting), opcodeToString[instr], d.code[i+1]))
			d.nesting += 1
			i += 1
		case asm.OpCodeJmpNZ:
			d.nesting -= 1
			result.WriteString(fmt.Sprintf("%s%s (%v)\n", strings.Repeat(" ", d.nesting), opcodeToString[instr], d.code[i+1]))
			i += 1
		}
	}

	return result.String()
}
