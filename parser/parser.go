package parser

import (
	"github.com/krekoten/8s/ast"
	"github.com/krekoten/8s/lexer"
)

var tokenToAstNode map[lexer.TokenType]ast.AstNode = map[lexer.TokenType]ast.AstNode{
	lexer.MoveRight: ast.AstNextCell{},
	lexer.MoveLeft:  ast.AstPrevCell{},
	lexer.Increment: ast.AstIncrement{},
	lexer.Decrement: ast.AstDecrement{},
	lexer.Input:     ast.AstInput{},
	lexer.Output:    ast.AstOutput{},
}

type Parser struct {
	tokens []lexer.Token
	pos    int
}

func New(tokens []lexer.Token) *Parser {
	return &Parser{tokens: tokens}
}

func (p *Parser) Parse() ast.AstStatements {
	stmt := ast.AstStatements{}

	for !p.end() {
		token := p.currentToken()
		if token.TokenType == lexer.Comment {
			p.next()
			continue
		}

		stmt.Statements = append(stmt.Statements, tokenToAstNode[token.TokenType])
		p.next()
	}

	return stmt
}

func (p *Parser) end() bool {
	return p.currentToken().TokenType == lexer.EndOfFile
}

func (p *Parser) currentToken() lexer.Token {
	return p.tokens[p.pos]
}

func (p *Parser) next() {
	p.pos += 1
}
