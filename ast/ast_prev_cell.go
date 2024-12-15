package ast

type AstPrevCell struct{}

func (node AstPrevCell) accept(v Visitor) {
	v.visit(node)
}

func (node AstPrevCell) String() string {
	return "AstPrevCell"
}
