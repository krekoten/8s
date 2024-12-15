package ast

type AstOutput struct{}

func (node AstOutput) ast() {}

func (node AstOutput) String() string {
	return "AstOutput"
}
