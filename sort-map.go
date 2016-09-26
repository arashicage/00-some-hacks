package main

import (
	"sort"

	"fmt"
)

func main() {
	m := map[string]int{
		"a": 6,
		"e": 1,
		"d": 2,
		"f": 4,
		"c": 4,
		"b": 3,
	}

	// 按 key 排序输出
	var keys []string

	for k, _ := range m {
		keys = append(keys, k)
	}

	// 按 key 排序输出
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, m[k])
	}

	// 按 val 排序输出
	// 新建一个map 对调 key 和 val
	// 照上面逻辑重新走一遍
	// 按照val 排序然后找对应的键再输出是不合适的，应为val 可能相同，这个时候取哪个呢

}
