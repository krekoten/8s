package ast

import "fmt"

type AstStatements struct {
	Statements []AstNode
}

func (node AstStatements) accept(v Visitor) {
	v.visit(node)
}

func (node AstStatements) String() string {
	return fmt.Sprintf("AstStatements{Statements: %s}", node.Statements)
}
