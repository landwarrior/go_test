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
	// success.Main()
}
