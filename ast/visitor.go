package ast

type Visitor interface {
	VisitDecrement(n AstDecrement)
	VisitIncrement(n AstIncrement)
	VisitNextCell(n AstNextCell)
	VisitPrevCell(n AstPrevCell)
	VisitInput(n AstInput)
	VisitOutput(n AstOutput)
	VisitStatements(n AstStatements)
	VisitLoop(n AstLoop)
	VisitExit(n AstExit)
}
