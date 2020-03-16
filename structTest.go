package main

import "fmt"

// 構造体型の作成
type myStruct struct {
	a int
	b string
	c string
}

func testStruct() {
	fmt.Println("構造体型のテスト ここから")
	// 初期化はそれぞれ代入しないといけない
	var test myStruct
	test.a = 1
	test.b = "test"
	test.c = "test"
	fmt.Println(test)
	fmt.Println("構造体型のテスト ここまで")
}
