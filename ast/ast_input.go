package ast

type AstInput struct{}

func (node AstInput) ast() {}

func (node AstInput) String() string {
	return "AstInput"
}
