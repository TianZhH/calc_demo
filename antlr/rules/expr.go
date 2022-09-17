package rules

import (
	"fmt"
)

/**
 * @Author: tianzhenhua.nathan@bytedance.com
 * @Date: 2022-09-17
 * @Desc:
 */

type ExprStatParent interface {
	AcceptExpr(expr *ExprStat) error
}

type ExprStat struct {
	Left     *ExprStat `json:"Left,omitempty"`
	Right    *ExprStat `json:"Right,omitempty"`
	Operator string    `json:"Operator,omitempty"`
	Num      *int64    `json:"Num,omitempty"`
}

func (e *ExprStat) Exec() (int64, error) {
	var err error
	if e.Num != nil {
		return *e.Num, nil
	}

	var left, right int64
	if e.Left != nil {
		if left, err = e.Left.Exec(); err != nil {
			return 0, err
		}
	}

	if e.Operator == "" {
		return left, nil
	}

	if e.Right != nil {
		if right, err = e.Right.Exec(); err != nil {
			return 0, err
		}
	}

	switch e.Operator {
	case "-":
		return left - right, nil
	case "+":
		return left + right, nil
	case "*":
		return left * right, nil
	case "/":
		return left / right, nil
	default:
		return 0, fmt.Errorf("unknown operator")
	}
}

func (e *ExprStat) AcceptExpr(expr *ExprStat) error {
	if e.Left == nil {
		e.Left = expr
		return nil
	}

	if e.Right == nil {
		e.Right = expr
		return nil
	}

	return fmt.Errorf("expr is defined")
}

func (e *ExprStat) AcceptNum(n int64) error {
	e.Num = &n
	return nil
}

func (e *ExprStat) AcceptOperator(operator string) error {
	e.Operator = operator
	return nil
}
