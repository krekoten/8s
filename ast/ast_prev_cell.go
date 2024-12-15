package ast

type AstPrevCell struct{}

func (node AstPrevCell) ast() {}

func (node AstPrevCell) String() string {
	return "AstPrevCell"
}
