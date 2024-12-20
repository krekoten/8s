package compiler

import (
	"github.com/krekoten/8s/asm"
	"github.com/krekoten/8s/ast"
)

type Compiler struct {
	// ip int
	ast   ast.AstStatements
	instr []asm.OpCode
	ip    int
}

func New(ast ast.AstStatements) *Compiler {
	return &Compiler{ast: ast, instr: make([]asm.OpCode, 0), ip: 0}
}

func (c *Compiler) emit(opc asm.OpCode) {
	c.instr = append(c.instr, opc)
	c.ip += 1
}

func (c *Compiler) emitWithValue(opc, value asm.OpCode) {
	c.instr = append(c.instr, opc, value)
	c.ip += 2
}

func (c *Compiler) updateWithValueAt(i int, value asm.OpCode) {
	c.instr[i] = value
}

func (c *Compiler) Compile() []asm.OpCode {
	c.ast.Accept(c)
	return c.instr
}

func (c *Compiler) VisitIncrement(n ast.AstIncrement) {
	c.emitWithValue(asm.OpcodeIncreaseBy, 1)
}

func (c *Compiler) VisitDecrement(n ast.AstDecrement) {
	c.emitWithValue(asm.OpCodeDecreaseBy, 1)
}

func (c *Compiler) VisitNextCell(n ast.AstNextCell) {
	c.emit(asm.OpCodeMoveRight)
}

func (c *Compiler) VisitPrevCell(n ast.AstPrevCell) {
	c.emit(asm.OpCodeMoveLeft)
}

func (c *Compiler) VisitInput(n ast.AstInput) {
	c.emit(asm.OpCodeInput)
}

func (c *Compiler) VisitOutput(n ast.AstOutput) {
	c.emit(asm.OpCodeOutput)
}

func (c *Compiler) VisitStatements(n ast.AstStatements) {
	for _, stmt := range n.Statements {
		stmt.Accept(c)
	}
}

func (c *Compiler) VisitExit(n ast.AstExit) {
	c.emit(asm.OpCodeNoOp)
}

func (c *Compiler) VisitLoop(n ast.AstLoop) {
	currentIp := c.ip

	c.emitWithValue(asm.OpCodeJmpZ, 0)
	c.VisitStatements(n.Statements)
	// next byte after JmpZ and address
	c.emitWithValue(asm.OpCodeJmpNZ, asm.OpCode(currentIp+2))

	endOfLoopIp := c.ip
	// patch JmpZ's address to next after ZmpNZ address cell
	c.updateWithValueAt(currentIp+1, asm.OpCode(endOfLoopIp))
}
