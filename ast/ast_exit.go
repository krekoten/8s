package ast

type AstExit struct{}

func (node AstExit) Accept(v Visitor) {
	v.VisitExit(node)
}

func (node AstExit) String() string {
	return "AstExit"
}
