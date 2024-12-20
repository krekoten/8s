package ast

import "fmt"

type AstStatements struct {
	Statements []AstNode
}

func (node AstStatements) Accept(v Visitor) {
	v.VisitStatements(node)
}

func (node AstStatements) String() string {
	return fmt.Sprintf("AstStatements{Statements: %s}", node.Statements)
}
