package ast

type AstNextCell struct{}

func (node AstNextCell) Accept(v Visitor) {
	v.VisitNextCell(node)
}

func (node AstNextCell) String() string {
	return "AstNextCell"
}
