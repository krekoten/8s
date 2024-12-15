package lexer

type Lexer struct {
	sourceCode      string
	currentPosition int
	tokens          []Token
}

var commandToTokenType map[byte]TokenType = map[byte]TokenType{
	'>': MoveRight,
	'<': MoveLeft,
	'+': Increment,
	'-': Decrement,
	',': Input,
	'.': Output,
	'[': LoopStart,
	']': LoopEnd,
}

func New(sourceCode string) *Lexer {
	return &Lexer{
		sourceCode:      sourceCode,
		currentPosition: 0,
		tokens:          make([]Token, 0),
	}
}

func (l *Lexer) endOfFile() bool {
	return l.currentPosition >= len(l.sourceCode)
}

func (l *Lexer) next() {
	l.currentPosition += 1
}

func (l *Lexer) currentTokenType() TokenType {
	command := l.sourceCode[l.currentPosition]

	return commandToTokenType[command]
}

func (l *Lexer) Tokenize() []Token {
	for !l.endOfFile() {
		l.tokens = append(l.tokens, NewToken(l.currentTokenType(), l.currentPosition))
		l.next()
	}

	l.tokens = append(l.tokens, NewToken(EndOfFile, l.currentPosition))

	return l.tokens
}
