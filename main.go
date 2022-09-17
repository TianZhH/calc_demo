package main

import (
	"calc_demo/builder"
	"fmt"
)

/**
 * @Author: tianzhenhua.nathan@bytedance.com
 * @Date: 2022-09-17
 * @Desc:
 */

func main() {
	rule, err := builder.BuildRule("1 + 2 * 3")
	if err != nil {
		println(fmt.Sprintf("err=%+v", err))
		return
	}

	//rule.Print()

	ret, err := rule.Exec()
	println(fmt.Sprintf("ret=%+v, err=%+v", ret, err))
}
