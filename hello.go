package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	// C# でいう partial class みたいに、別ファイルの関数を実行可能みたいね
	// ただし、 go build でファイル指定なしでまとめてビルドしないとダメみたい
	testStruct()
	pointerTest()
	testArray()
	testSlice()
	testMapping()
	a := 1
	b := "hello"
	c := b + string(a)
	fmt.Printf("%v\n", c)
	d := fmt.Sprintf("%v%v", b, a)
	fmt.Println(d)
	testCondition()
	testRoop()
	testRoop2("a", "b", "c", "d")
	fmt.Println(normalFunc(3, 5))
	// success.Main()
}

func normalFunc(a int, b int) (int, int) {
	fmt.Println("標準的な関数のテスト ここから")
	// 普通の関数
	a2 := a * 2
	b2 := b * 2
	fmt.Println("標準的な関数のテスト ここまで（reeturn があるから）")
	return a2, b2
}
