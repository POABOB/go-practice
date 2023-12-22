package utils

import (
	"sort"

	"golang.org/x/exp/constraints"
)

type OrderMap[K constraints.Ordered, V any] struct {
	// 儲存升序的 key 值
	keys []K
	// 原本的 map
	originMap map[K]V
}

// NewOrderMap 實例 OrderMap
func NewOrderMap[K constraints.Ordered, V any](m map[K]V) *OrderMap[K, V] {
	// 排序 keys 值
	k := make([]K, 0, len(m))
	for key := range m {
		k = append(k, key)
	}

	sort.Slice(k, func(i, j int) bool {
		return k[i] < k[j]
	})

	return &OrderMap[K, V]{
		keys:      k,
		originMap: m,
	}
}

// Get 獲取 Value，相當於 v, ok := o.originMap[key]
func (o *OrderMap[K, V]) Get(key K) (V, bool) {
	v, ok := o.originMap[key]
	return v, ok
}

// GetAscKeys 獲取升序的 Key 值
func (o *OrderMap[K, V]) GetAscKeys() []K {
	return o.keys
}

// GetDescKeys 獲取降序的 Key 值
func (o *OrderMap[K, V]) GetDescKeys() []K {
	descKeys := make([]K, 0, len(o.keys))
	for i := len(o.keys) - 1; i >= 0; i-- {
		descKeys = append(descKeys, o.keys[i])
	}
	return descKeys
}

// IterateInAsc 升序遍歷
func (o *OrderMap[K, V]) IterateInAsc(fn func(k K, v V)) {
	for _, k := range o.keys {
		fn(k, o.originMap[k])
	}
}

// IterateInDesc 降序遍歷
func (o *OrderMap[K, V]) IterateInDesc(fn func(k K, v V)) {
	descKeys := o.GetDescKeys()
	for _, k := range descKeys {
		fn(k, o.originMap[k])
	}
}

// Append 新增 Map 值，如果 key 重複就蓋過去
func (o *OrderMap[K, V]) Append(key K, value V) {
	// 重複的 key
	if _, ok := o.originMap[key]; ok {
		o.originMap[key] = value
		return
	}

	o.originMap[key] = value
	index := o.binarySearch(key) + 1
	if index == len(o.keys)-1 { // 要插入的 key 位置在最後面
		o.keys = append(o.keys, key)
	} else { // 要插入的 key 位置在 [0, len(o.keys) - 1)
		temp := make([]K, len(o.keys)-index)
		copy(temp, o.keys[index:])
		o.keys = o.keys[:index]
		o.keys = append(o.keys, key)
		o.keys = append(o.keys, temp...)
	}
}

// Delete 刪除 Map 值，有值就刪除，沒值就忽略
func (o *OrderMap[K, V]) Delete(key K) {
	if _, ok := o.originMap[key]; ok {
		delete(o.originMap, key)
		index := o.binarySearch(key)
		temp := o.keys[index+1:]
		o.keys = o.keys[:index]
		o.keys = append(o.keys, temp...)
	}
}

// binarySearch 搜尋 upper bound 的 index 位置
func (o *OrderMap[K, V]) binarySearch(key K) int {
	left, right := 0, len(o.keys)-1
	for right > left {
		mid := left + (right-left)>>1

		if key < o.keys[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right - 1
}
