package ast

type Visitor interface {
	visit(n AstNode)
}

type AstNode interface {
	accept(v Visitor)
}
