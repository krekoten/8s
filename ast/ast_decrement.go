package ast

type AstDecrement struct{}

func (node AstDecrement) ast() {}

func (node AstDecrement) String() string {
	return "AstDecrement"
}
