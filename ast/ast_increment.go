package ast

type AstIncrement struct{}

func (node AstIncrement) ast() {}

func (node AstIncrement) String() string {
	return "AstIncrement"
}
