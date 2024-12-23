package parser

import (
	"github.com/krekoten/8s/ast"
	"github.com/krekoten/8s/lexer"
)

func defaultHandler(node ast.AstNode) func() ast.AstNode {
	return func() ast.AstNode {
		return node
	}
}

var tokenToAstNode map[lexer.TokenType]func() ast.AstNode = map[lexer.TokenType]func() ast.AstNode{
	lexer.MoveRight: defaultHandler(ast.AstNextCell{}),
	lexer.MoveLeft:  defaultHandler(ast.AstPrevCell{}),
	lexer.Input:     defaultHandler(ast.AstInput{}),
	lexer.Output:    defaultHandler(ast.AstOutput{}),
	lexer.Increment: func() ast.AstNode {
		return ast.AstIncrement{Value: 1}
	},
	lexer.Decrement: func() ast.AstNode {
		return ast.AstDecrement{Value: 1}
	},
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

		if token.TokenType == lexer.LoopStart {
			loop := p.parseLoop()
			stmt.Statements = append(stmt.Statements, loop)
		} else {
			stmt.Statements = append(stmt.Statements, tokenToAstNode[token.TokenType]())
		}

		p.next()
	}

	if p.end() {
		stmt.Statements = append(stmt.Statements, ast.AstExit{})
	}

	return stmt
}

func (p *Parser) parseLoop() ast.AstLoop {
	stmt := ast.AstStatements{}

	p.next() // skip loop start command

	for !p.end() && p.currentToken().TokenType != lexer.LoopEnd {
		token := p.currentToken()
		if token.TokenType == lexer.Comment {
			p.next()
			continue
		}

		if token.TokenType == lexer.LoopStart {
			nestedLoop := p.parseLoop()
			stmt.Statements = append(stmt.Statements, nestedLoop)
		} else {
			stmt.Statements = append(stmt.Statements, tokenToAstNode[token.TokenType]())
		}

		p.next()
	}

	return ast.AstLoop{Statements: stmt}
}

func (p *Parser) end() bool {
	return p.pos >= len(p.tokens) || p.currentToken().TokenType == lexer.EndOfFile
}

func (p *Parser) currentToken() lexer.Token {
	return p.tokens[p.pos]
}

func (p *Parser) next() {
	p.pos += 1
}
