package ast

import "fmt"

type AstIncrement struct {
	Value byte
}

func (node AstIncrement) Accept(v Visitor) {
	v.VisitIncrement(node)
}

func (node AstIncrement) String() string {
	return fmt.Sprintf("AstIncrement{Value: %v}", node.Value)
}
