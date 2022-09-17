package rules

/**
 * @Author: tianzhenhua.nathan@bytedance.com
 * @Date: 2022-09-17
 * @Desc:
 */

type OperatorParent interface {
	AcceptOperator(operator string) error
}
