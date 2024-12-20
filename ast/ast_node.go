package ast

type AstNode interface {
	Accept(v Visitor)
}
