package compiler_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/krekoten/8s/asm"
	"github.com/krekoten/8s/ast"
	"github.com/krekoten/8s/compiler"
)

func TestCompilation(t *testing.T) {
	ast := ast.AstStatements{
		Statements: []ast.AstNode{
			ast.AstNextCell{},
			ast.AstPrevCell{},
			ast.AstIncrement{},
			ast.AstDecrement{},
			ast.AstInput{},
			ast.AstOutput{},
			ast.AstLoop{
				Statements: ast.AstStatements{
					Statements: []ast.AstNode{
						ast.AstIncrement{},
					},
				},
			},
			ast.AstExit{},
		},
	}

	c := compiler.New(ast)
	actual := c.Compile()

	expected := []asm.OpCode{
		// OpCode									IP
		asm.OpCodeMoveRight,  // 	 0
		asm.OpCodeMoveLeft,   //	 1
		asm.OpcodeIncreaseBy, //	 2
		1,                    //	 3
		asm.OpCodeDecreaseBy, //	 4
		1,                    //	 5
		asm.OpCodeInput,      //	 6
		asm.OpCodeOutput,     //	 7
		asm.OpCodeJmpZ,       //   8
		14,                   //	 9
		asm.OpcodeIncreaseBy, //	10
		1,                    //	11
		asm.OpCodeJmpNZ,      //	12
		10,                   //	13
		asm.OpCodeNoOp,       //  14
	}

	if !cmp.Equal(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
