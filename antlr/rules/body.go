package rules

import "fmt"

/**
 * @Author: tianzhenhua.nathan@bytedance.com
 * @Date: 2022-09-17
 * @Desc:
 */

type BodyStat struct {
	Expr *ExprStat `json:"Expr,omitempty"`
}

func (b *BodyStat) AcceptExpr(expr *ExprStat) error {
	b.Expr = expr
	return nil
}

func (b *BodyStat) Exec() (int64, error) {
	if b.Expr != nil {
		return b.Expr.Exec()
	}

	return 0, fmt.Errorf("expr is not definied in body")
}
