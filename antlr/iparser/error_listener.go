package iparser

import (
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

/**
 * @Author: tianzhenhua.nathan@bytedance.com
 * @Date: 2022-09-17
 * @Desc:
 */

const (
	SyntaxErrFormat = "[%d:%d]:%s"
)

func packageDslErrMsg(line, column int, msg string) string {
	return fmt.Sprintf(SyntaxErrFormat, line, column, msg)
}

type CalcErrorListener struct {
	antlr.ErrorListener
	GrammarErrors []string
}

func NewCalcErrorListener() *CalcErrorListener {
	return &CalcErrorListener{}
}

func (g *CalcErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	g.GrammarErrors = append(g.GrammarErrors, packageDslErrMsg(line, column, msg))
}

func (g *CalcErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	//println(fmt.Sprintf("ReportAmbiguity:"))
}
func (g *CalcErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	//println(fmt.Sprintf("g=%p, ReportAttemptingFullContext:", g))
}
func (g *CalcErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	//println(fmt.Sprintf("ReportContextSensitivity:"))
}

// GetSyntaxErr tianzhenhua@2021/12/30:
func (g *CalcErrorListener) GetSyntaxErr() error {
	if len(g.GrammarErrors) > 0 {
		return NewSyntaxErr(strings.Join(g.GrammarErrors, "\n"))
	}

	return nil
}
