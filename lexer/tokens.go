package lexer

type TokenType int

const (
	Comment TokenType = iota
	MoveRight
	MoveLeft
	Increment
	Decrement
	Output
	Input
	LoopStart
	LoopEnd
	EndOfFile
)

var tokenToString map[TokenType]string = map[TokenType]string{
	Comment:   "Comment",
	MoveRight: "MoveRight",
	MoveLeft:  "MoveLeft",
	Increment: "Increment",
	Decrement: "Decrement",
	Output:    "Output",
	Input:     "Input",
	LoopStart: "LoopStart",
	LoopEnd:   "LoopEnd",
	EndOfFile: "EndOfFile",
}

type Token struct {
	TokenType TokenType
	Pos       int
}

func NewToken(tokenType TokenType, pos int) Token {
	return Token{TokenType: tokenType, Pos: pos}
}

func (t Token) String() string {
	return tokenToString[t.TokenType]
}
