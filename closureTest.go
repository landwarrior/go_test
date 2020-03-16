package main

import "fmt"

// Go 言語では JavaScript のようにクロージャが使える
func testClosure() {
	// defer を付けると最後に実行される
	defer fmt.Println("defer!")
	fmt.Println("クロージャのテスト ここから")
	var fnc = f()
	fmt.Println(fnc())
	fmt.Println(fnc())
	fmt.Println(fnc())
	fmt.Println("クロージャのテスト ここまで")
}

func f() func() int {
	alpha := 0

	// さすがに1回しか出なかったわ
	defer fmt.Println("クロージャ！")
	// ローカル変数を参照する関数を返却する
	return func() int {
		alpha++
		return alpha
	}
}
