package iparser

import (
	"calc_demo/antlr/calc_parse/parser"
	"calc_demo/antlr/rules"
	"calc_demo/utils"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

/**
 * @Author: tianzhenhua.nathan@bytedance.com
 * @Date: 2022-09-17
 * @Desc:
 */

type CalcListener struct {
	parser.BaseCalcListener

	Stack       *utils.Stack
	ParseErrors []string
}

// NewCaleParserListener tianzhenhua@2021/9/30:
func NewCaleParserListener() *CalcListener {
	return &CalcListener{
		Stack: utils.NewStack(),
	}
}

func (g *CalcListener) AddError(errMsg string) {
	g.ParseErrors = append(g.ParseErrors, errMsg)
}

func (g *CalcListener) HasError() bool {
	return len(g.ParseErrors) > 0
}

func (g *CalcListener) GetSyntaxErr() error {
	if len(g.ParseErrors) > 0 {
		return NewSyntaxErr(strings.Join(g.ParseErrors, "\n"))
	}

	return nil
}

//// EnterBody is called when entering the body production.
//	EnterBody(c *BodyContext)
//
//	// EnterExpr is called when entering the expr production.
//	EnterExpr(c *ExprContext)
//
//	// EnterOperator is called when entering the operator production.
//	EnterOperator(c *OperatorContext)
//
//	// EnterNum is called when entering the num production.
//	EnterNum(c *NumContext)
//
//	// ExitBody is called when exiting the body production.
//	ExitBody(c *BodyContext)
//
//	// ExitExpr is called when exiting the expr production.
//	ExitExpr(c *ExprContext)
//
//	// ExitOperator is called when exiting the operator production.
//	ExitOperator(c *OperatorContext)
//
//	// ExitNum is called when exiting the num production.
//	ExitNum(c *NumContext)

// NewSyntaxErr tianzhenhua@2021/12/30:
func NewSyntaxErr(syntaxDesc string) error {
	return errors.New("syntax error:\n" + syntaxDesc)
}

// EnterBody is called when production ruleBody is entered.
func (g *CalcListener) EnterBody(ctx *parser.BodyContext) {
	println(fmt.Sprintf("EnterBody"))

	if g.HasError() {
		return
	}

	g.Stack.Push(&rules.BodyStat{})
	return
}

// ExitBody is called when production ruleBody is exited.
func (g *CalcListener) ExitBody(ctx *parser.BodyContext) {
	println(fmt.Sprintf("ExitBody"))
}

func (g *CalcListener) EnterExpr(ctx *parser.ExprContext) {
	println(fmt.Sprintf("EnterExpr"))

	if g.HasError() {
		return
	}

	g.Stack.Push(&rules.ExprStat{})

	return
}
func (g *CalcListener) ExitExpr(ctx *parser.ExprContext) {
	println(fmt.Sprintf("ExitExpr"))

	if g.HasError() {
		return
	}

	exprRule := g.Stack.Pop().(*rules.ExprStat)
	exprRuleParent := g.Stack.Peek().(rules.ExprStatParent)

	if err := exprRuleParent.AcceptExpr(exprRule); err != nil {
		g.AddError(packageDslErrMsg(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), err.Error()))
		return
	}

	return
}

// EnterMd_operator is called when entering the md_operator production.
func (g *CalcListener) EnterMd_operator(ctx *parser.Md_operatorContext) {
	println(fmt.Sprintf("EnterMd_operator"))
}

func (g *CalcListener) ExitMd_operator(ctx *parser.Md_operatorContext) {
	println(fmt.Sprintf("ExitMd_operator"))

	if g.HasError() {
		return
	}

	operator := ctx.GetText()

	integerHolder := g.Stack.Peek().(rules.OperatorParent)
	err := integerHolder.AcceptOperator(operator)
	if err != nil {
		g.AddError(packageDslErrMsg(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), err.Error()))
		return
	}
}

// EnterAs_operator is called when entering the as_operator production.
func (g *CalcListener) EnterAs_operator(ctx *parser.As_operatorContext) {
	println(fmt.Sprintf("EnterAs_operator"))
}

func (g *CalcListener) ExitAs_operator(ctx *parser.As_operatorContext) {
	println(fmt.Sprintf("ExitAs_operator"))

	if g.HasError() {
		return
	}

	operator := ctx.GetText()

	integerHolder := g.Stack.Peek().(rules.OperatorParent)
	err := integerHolder.AcceptOperator(operator)
	if err != nil {
		g.AddError(packageDslErrMsg(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), err.Error()))
		return
	}
}

func (g *CalcListener) EnterNum(ctx *parser.NumContext) {
	println(fmt.Sprintf("EnterNum"))
}

func (g *CalcListener) ExitNum(ctx *parser.NumContext) {
	println(fmt.Sprintf("ExitNum"))

	if g.HasError() {
		return
	}

	i, err := strconv.ParseInt(ctx.GetText(), 10, 64)
	if err != nil {
		g.AddError(packageDslErrMsg(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), fmt.Sprintf("parse int64 failed, text=%s, err=%+v", ctx.GetText(), err)))
		return
	}

	integerHolder := g.Stack.Peek().(rules.NumParent)
	err = integerHolder.AcceptNum(i)
	if err != nil {
		g.AddError(packageDslErrMsg(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), err.Error()))
		return
	}
}
