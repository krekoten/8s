package ast

type AstPrevCell struct{}

func (node AstPrevCell) Accept(v Visitor) {
	v.VisitPrevCell(node)
}

func (node AstPrevCell) String() string {
	return "AstPrevCell"
}
