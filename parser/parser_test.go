package parser_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/krekoten/8s/ast"
	"github.com/krekoten/8s/lexer"
	"github.com/krekoten/8s/parser"
)

func TestParsingSimpleCommands(t *testing.T) {
	tokens := []lexer.Token{
		lexer.NewToken(lexer.MoveRight, 0),
		lexer.NewToken(lexer.MoveLeft, 1),
		lexer.NewToken(lexer.Increment, 2),
		lexer.NewToken(lexer.Decrement, 3),
		lexer.NewToken(lexer.Input, 4),
		lexer.NewToken(lexer.Output, 5),
		lexer.NewToken(lexer.EndOfFile, 6),
	}
	p := parser.New(tokens)

	expected := ast.AstStatements{
		Statements: []ast.AstNode{
			ast.AstNextCell{},
			ast.AstPrevCell{},
			ast.AstIncrement{Value: 1},
			ast.AstDecrement{Value: 1},
			ast.AstInput{},
			ast.AstOutput{},
			ast.AstExit{},
		},
	}

	actual := p.Parse()

	if !cmp.Equal(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestParsingLoops(t *testing.T) {
	tokens := []lexer.Token{
		lexer.NewToken(lexer.LoopStart, 0),
		lexer.NewToken(lexer.Increment, 1),
		lexer.NewToken(lexer.LoopEnd, 2),
		lexer.NewToken(lexer.EndOfFile, 3),
	}
	p := parser.New(tokens)

	expected := ast.AstStatements{
		Statements: []ast.AstNode{
			ast.AstLoop{
				Statements: ast.AstStatements{
					Statements: []ast.AstNode{
						ast.AstIncrement{Value: 1},
					},
				},
			},
			ast.AstExit{},
		},
	}

	actual := p.Parse()

	if !cmp.Equal(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestParsingNestedLoops(t *testing.T) {
	tokens := []lexer.Token{
		lexer.NewToken(lexer.LoopStart, 0),
		lexer.NewToken(lexer.Increment, 1),
		lexer.NewToken(lexer.LoopStart, 2),
		lexer.NewToken(lexer.Decrement, 3),
		lexer.NewToken(lexer.LoopEnd, 4),
		lexer.NewToken(lexer.LoopEnd, 5),
		lexer.NewToken(lexer.EndOfFile, 6),
	}
	p := parser.New(tokens)

	expected := ast.AstStatements{
		Statements: []ast.AstNode{
			ast.AstLoop{
				Statements: ast.AstStatements{
					Statements: []ast.AstNode{
						ast.AstIncrement{Value: 1},
						ast.AstLoop{
							Statements: ast.AstStatements{
								Statements: []ast.AstNode{
									ast.AstDecrement{Value: 1},
								},
							},
						},
					},
				},
			},
			ast.AstExit{},
		},
	}

	actual := p.Parse()

	if !cmp.Equal(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
