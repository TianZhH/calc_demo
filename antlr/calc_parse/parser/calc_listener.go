// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Calc

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// CalcListener is a complete listener for a parse tree produced by CalcParser.
type CalcListener interface {
	antlr.ParseTreeListener

	// EnterBody is called when entering the body production.
	EnterBody(c *BodyContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterMd_operator is called when entering the md_operator production.
	EnterMd_operator(c *Md_operatorContext)

	// EnterAs_operator is called when entering the as_operator production.
	EnterAs_operator(c *As_operatorContext)

	// EnterNum is called when entering the num production.
	EnterNum(c *NumContext)

	// ExitBody is called when exiting the body production.
	ExitBody(c *BodyContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitMd_operator is called when exiting the md_operator production.
	ExitMd_operator(c *Md_operatorContext)

	// ExitAs_operator is called when exiting the as_operator production.
	ExitAs_operator(c *As_operatorContext)

	// ExitNum is called when exiting the num production.
	ExitNum(c *NumContext)
}
