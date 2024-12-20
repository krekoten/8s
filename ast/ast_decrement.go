package ast

import "fmt"

type AstDecrement struct {
	Value byte
}

func (node AstDecrement) Accept(v Visitor) {
	v.VisitDecrement(node)
}

func (node AstDecrement) String() string {
	return fmt.Sprintf("AstDecrement{Value: %v}", node.Value)
}
