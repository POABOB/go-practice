package main

import (
	"fmt"
	"github.com/POABOB/go-practice/order-map/v2/utils"
)

func main() {
	m := map[string]int{
		"Alice": 1,
		"Bob":   2,
		"Cindy": 3,
		"Duke":  4,
	}
	orderMap := utils.NewOrderMap[string, int](m)

	fmt.Println(orderMap.GetAscKeys())
	fmt.Println(orderMap.GetDescKeys())

	fmt.Println("====================")

	fmt.Println(orderMap.Get("Bob"))
	fmt.Println(orderMap.Get("Bob1"))

	orderMap.Delete("Bob")

	orderMap.Append("Eric", 5)
	orderMap.Append("Bobo", 6)
	orderMap.Append("Zack", 7)

	fmt.Println("====================")

	fmt.Println(orderMap.GetAscKeys())
	fmt.Println(orderMap.GetDescKeys())

	fmt.Println("====================")

	orderMap.IterateInAsc(func(k string, v int) {
		fmt.Printf("key=%v, value=%v\n", k, v)
	})

	fmt.Println("====================")

	orderMap.IterateInDesc(func(k string, v int) {
		fmt.Printf("key=%v, value=%v\n", k, v)
	})

}
