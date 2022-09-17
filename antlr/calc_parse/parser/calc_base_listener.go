// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Calc

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseCalcListener is a complete listener for a parse tree produced by CalcParser.
type BaseCalcListener struct{}

var _ CalcListener = &BaseCalcListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCalcListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCalcListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCalcListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCalcListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterBody is called when production body is entered.
func (s *BaseCalcListener) EnterBody(ctx *BodyContext) {}

// ExitBody is called when production body is exited.
func (s *BaseCalcListener) ExitBody(ctx *BodyContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseCalcListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseCalcListener) ExitExpr(ctx *ExprContext) {}

// EnterMd_operator is called when production md_operator is entered.
func (s *BaseCalcListener) EnterMd_operator(ctx *Md_operatorContext) {}

// ExitMd_operator is called when production md_operator is exited.
func (s *BaseCalcListener) ExitMd_operator(ctx *Md_operatorContext) {}

// EnterAs_operator is called when production as_operator is entered.
func (s *BaseCalcListener) EnterAs_operator(ctx *As_operatorContext) {}

// ExitAs_operator is called when production as_operator is exited.
func (s *BaseCalcListener) ExitAs_operator(ctx *As_operatorContext) {}

// EnterNum is called when production num is entered.
func (s *BaseCalcListener) EnterNum(ctx *NumContext) {}

// ExitNum is called when production num is exited.
func (s *BaseCalcListener) ExitNum(ctx *NumContext) {}
