package main

import "fmt"

// 関数は引数や戻り値にすることもできるらしい
func testStoreFunc() {
	fmt.Println("関数の代入テスト ここから")
	// 関数型の変数を定義
	var fnc func(int, int) int

	// 関数の代入
	fnc = add
	fmt.Println(fnc(2, 1))
	// 関数の代入
	fnc = mul
	fmt.Println(fnc(2, 1))
	fmt.Println("関数の代入テスト ここまで")
}

func add(a int, b int) int {
	return a + b
}

func mul(a int, b int) int {
	return a * b
}
