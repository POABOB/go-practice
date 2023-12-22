package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{
		"Alice": 1,
		"Bob":   2,
		"Cindy": 3,
		"Duke":  4,
	}

	// keys 存放順序的 key
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	// 排序
	sort.Strings(keys)

	// 根據 keys 順序來輸出 map 的值
	for _, k := range keys {
		fmt.Printf("key=%s, value=%d\n", k, m[k])
	}
}
