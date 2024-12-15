package ast

type AstOutput struct{}

func (node AstOutput) accept(v Visitor) {
	v.visit(node)
}

func (node AstOutput) String() string {
	return "AstOutput"
}
