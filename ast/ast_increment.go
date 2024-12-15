package ast

type AstIncrement struct{}

func (node AstIncrement) accept(v Visitor) {
	v.visit(node)
}

func (node AstIncrement) String() string {
	return "AstIncrement"
}
