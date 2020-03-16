package main

import (
	"fmt"
)

func pointerTest() {
	fmt.Println("ポインターのテスト ここから")
	// ポインターの宣言
	var p *int
	var test *int
	n := 1
	p = &n
	a := 10
	test = &a
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(test)
	fmt.Println(*test)
	fmt.Println("ポインターのテスト ここまで")
}
