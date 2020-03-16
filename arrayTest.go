package main

import (
	"fmt"
)

func testArray() {
	fmt.Println("配列のテスト ここから")
	test := [...]int{1, 2, 3}
	test2 := [...]int{1: 2, 3: 4}
	test3 := [3]string{}
	test4 := [...]int{1, 2, 3}
	test5 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(test)
	fmt.Println(test2)
	fmt.Printf("%T, %v\n", test3, test3)
	test2[0] = -5
	fmt.Println(test2)
	fmt.Println(test == test4)
	fmt.Printf("type:%T value:%v len:%d cap:%d\n", test5, test5, len(test5), cap(test5))
	fmt.Println(test5[1:6])
	fmt.Println("配列のテスト ここまで")
}
