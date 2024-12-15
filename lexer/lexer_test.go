package lexer_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/krekoten/8s/lexer"
)

func TestMoveRightTokenized(t *testing.T) {
	sourceCode := ">"
	expected := []lexer.Token{lexer.NewToken(lexer.MoveRight, 0), lexer.NewToken(lexer.EndOfFile, 1)}

	l := lexer.New(sourceCode)
	actual := l.Tokenize()

	if !cmp.Equal(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestMoveLeftTokenized(t *testing.T) {
	sourceCode := "<"
	expected := []lexer.Token{lexer.NewToken(lexer.MoveLeft, 0), lexer.NewToken(lexer.EndOfFile, 1)}

	l := lexer.New(sourceCode)
	actual := l.Tokenize()

	if !cmp.Equal(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestIncrementTokenized(t *testing.T) {
	sourceCode := "+"
	expected := []lexer.Token{lexer.NewToken(lexer.Increment, 0), lexer.NewToken(lexer.EndOfFile, 1)}

	l := lexer.New(sourceCode)
	actual := l.Tokenize()

	if !cmp.Equal(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestDecrementTokenized(t *testing.T) {
	sourceCode := "-"
	expected := []lexer.Token{lexer.NewToken(lexer.Decrement, 0), lexer.NewToken(lexer.EndOfFile, 1)}

	l := lexer.New(sourceCode)
	actual := l.Tokenize()

	if !cmp.Equal(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestInputTokenized(t *testing.T) {
	sourceCode := ","
	expected := []lexer.Token{lexer.NewToken(lexer.Input, 0), lexer.NewToken(lexer.EndOfFile, 1)}

	l := lexer.New(sourceCode)
	actual := l.Tokenize()

	if !cmp.Equal(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestOutputTokenized(t *testing.T) {
	sourceCode := "."
	expected := []lexer.Token{lexer.NewToken(lexer.Output, 0), lexer.NewToken(lexer.EndOfFile, 1)}

	l := lexer.New(sourceCode)
	actual := l.Tokenize()

	if !cmp.Equal(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestLoopStartTokenized(t *testing.T) {
	sourceCode := "["
	expected := []lexer.Token{lexer.NewToken(lexer.LoopStart, 0), lexer.NewToken(lexer.EndOfFile, 1)}

	l := lexer.New(sourceCode)
	actual := l.Tokenize()

	if !cmp.Equal(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestLoopEndTokenized(t *testing.T) {
	sourceCode := "]"
	expected := []lexer.Token{lexer.NewToken(lexer.LoopEnd, 0), lexer.NewToken(lexer.EndOfFile, 1)}

	l := lexer.New(sourceCode)
	actual := l.Tokenize()

	if !cmp.Equal(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestCommentTokenized(t *testing.T) {
	sourceCode := "a"
	expected := []lexer.Token{lexer.NewToken(lexer.Comment, 0), lexer.NewToken(lexer.EndOfFile, 1)}

	l := lexer.New(sourceCode)
	actual := l.Tokenize()

	if !cmp.Equal(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestMultipleCommentsTokenized(t *testing.T) {
	sourceCode := "abc some code"
	expected := []lexer.Token{lexer.NewToken(lexer.Comment, 0), lexer.NewToken(lexer.EndOfFile, len(sourceCode))}

	l := lexer.New(sourceCode)
	actual := l.Tokenize()

	if !cmp.Equal(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
