package rules

/**
 * @Author: tianzhenhua.nathan@bytedance.com
 * @Date: 2022-09-17
 * @Desc:
 */

type NumParent interface {
	AcceptNum(i int64) error
}
