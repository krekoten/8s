package vm

import (
	"fmt"

	"github.com/krekoten/8s/asm"
)

type vm struct {
	data   [30_000]byte
	code   []asm.OpCode
	ip, dp int
}

func New(code []asm.OpCode) *vm {
	return &vm{code: code}
}

func (vm *vm) Run() {
	for !vm.exit() {
		switch vm.currentInstruction() {
		case asm.OpCodeNoOp:
			vm.ip += 1
		case asm.OpCodeMoveRight:
			vm.dp += 1
			vm.ip += 1
		case asm.OpCodeMoveLeft:
			vm.dp -= 1
			vm.ip += 1
		case asm.OpCodeInput:
			var char byte
			fmt.Scanf("%c", &char)
			vm.setCurrentData(char)
			vm.ip += 1
		case asm.OpCodeOutput:
			fmt.Printf("%c", vm.currentData())
			vm.ip += 1
		case asm.OpcodeIncreaseBy:
			vm.ip += 1
			incr := vm.currentInstruction()
			vm.setCurrentData(vm.currentData() + byte(incr))
			vm.ip += 1
		case asm.OpCodeDecreaseBy:
			vm.ip += 1
			incr := vm.currentInstruction()
			vm.setCurrentData(vm.currentData() - byte(incr))
			vm.ip += 1
		case asm.OpCodeJmpZ:
			vm.ip += 1
			page := int(vm.currentInstruction())
			vm.ip += 1
			offset := int(vm.currentInstruction())
			addr := page*256 + offset
			if vm.currentData() == 0 {
				vm.ip = int(addr)
			} else {
				vm.ip += 1
			}
		case asm.OpCodeJmpNZ:
			vm.ip += 1
			page := int(vm.currentInstruction())
			vm.ip += 1
			offset := int(vm.currentInstruction())
			addr := page*256 + offset
			if vm.currentData() != 0 {
				vm.ip = int(addr)
			} else {
				vm.ip += 1
			}
		}
	}
}

func (vm *vm) exit() bool {
	return vm.ip >= len(vm.code)
}

func (vm *vm) currentInstruction() asm.OpCode {
	return vm.code[vm.ip]
}

func (vm *vm) currentData() byte {
	return vm.data[vm.dp]
}

func (vm *vm) setCurrentData(value byte) {
	vm.data[vm.dp] = value
}
