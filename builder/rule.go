package builder

import (
	"calc_demo/antlr/calc_parse/parser"
	"calc_demo/antlr/iparser"
	"calc_demo/antlr/rules"
	"calc_demo/utils"
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

/**
 * @Author: tianzhenhua.nathan@bytedance.com
 * @Date: 2022-09-17
 * @Desc:
 */

// Rule 策略
type Rule struct {
	ruleBody *rules.BodyStat
}

// BuildRule tianzhenhua@2021/10/13: 创建 rule
func BuildRule(rule string) (*Rule, error) {
	// Setup the input
	is := antlr.NewInputStream(rule)

	// Create the Lexer
	lexer := parser.NewCalcLexer(is)
	lexerErrListener := iparser.NewCalcErrorListener()
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexerErrListener)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewCalcParser(stream)
	grammarErrListener := iparser.NewCalcErrorListener()
	p.RemoveErrorListeners()
	p.AddErrorListener(grammarErrListener)

	// Finally parse the expression (by walking the tree)
	listener := iparser.NewCaleParserListener()
	body := p.Body()

	if err := lexerErrListener.GetSyntaxErr(); err != nil {
		return nil, err
	}

	if err := grammarErrListener.GetSyntaxErr(); err != nil {
		return nil, err
	}

	antlr.ParseTreeWalkerDefault.Walk(listener, body)
	if err := listener.GetSyntaxErr(); err != nil {
		return nil, err
	}

	return &Rule{ruleBody: listener.Stack.Peek().(*rules.BodyStat)}, nil
}

func (r *Rule) Exec() (int64, error) {
	if r == nil {
		return 0, fmt.Errorf("rule is nil")
	}

	return r.ruleBody.Exec()
}

func (r *Rule) Print() {
	println(fmt.Sprintf("body=%s", utils.InterfaceToJson(r.ruleBody)))
}