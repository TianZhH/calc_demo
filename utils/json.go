package utils

import "encoding/json"

/**
 * @Author: tianzhenhua.nathan@bytedance.com
 * @Date: 2022-09-17
 * @Desc:
 */

func InterfaceToJson(i interface{}) string {
	bs, _ := json.Marshal(i)
	return string(bs)
}
