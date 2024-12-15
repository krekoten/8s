package ast

type AstDecrement struct{}

func (node AstDecrement) accept(v Visitor) {
	v.visit(node)
}

func (node AstDecrement) String() string {
	return "AstDecrement"
}
