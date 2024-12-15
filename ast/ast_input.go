package ast

type AstInput struct{}

func (node AstInput) accept(v Visitor) {
	v.visit(node)
}

func (node AstInput) String() string {
	return "AstInput"
}
