package ast

type AstInput struct{}

func (node AstInput) Accept(v Visitor) {
	v.VisitInput(node)
}

func (node AstInput) String() string {
	return "AstInput"
}
