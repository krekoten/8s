package ast

type AstNextCell struct{}

func (node AstNextCell) ast() {}

func (node AstNextCell) String() string {
	return "AstNextCell"
}
