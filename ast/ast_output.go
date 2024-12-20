package ast

type AstOutput struct{}

func (node AstOutput) Accept(v Visitor) {
	v.VisitOutput(node)
}

func (node AstOutput) String() string {
	return "AstOutput"
}
