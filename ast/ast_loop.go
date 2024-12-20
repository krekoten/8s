package ast

import "fmt"

type AstLoop struct {
	Statements AstStatements
}

func (node AstLoop) Accept(v Visitor) {
	v.VisitLoop(node)
}

func (node AstLoop) String() string {
	return fmt.Sprintf("AstLoop{Statements: %s}", node.Statements)
}
