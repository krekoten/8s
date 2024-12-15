package ast

type AstNextCell struct{}

func (node AstNextCell) accept(v Visitor) {
	v.visit(node)
}

func (node AstNextCell) String() string {
	return "AstNextCell"
}
