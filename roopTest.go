package main

import "fmt"

func testRoop() {
	fmt.Println("ループのテスト ここから")
	for i := 0; i < 4; i++ {
		fmt.Println(i)
	}
	// foreach
	for i, t := range []int{4, 5, 6, 7} {
		fmt.Println(i, t)
	}
	for i, t := range [4]int{} {
		// python の range みたいに、回したい回数だけ決めてる時に使えそう
		fmt.Println(i, t)
	}
	// while
	w := 0
	for w < 2 {
		fmt.Println(w)
		w++
	}
	fmt.Println("ループのテスト ここまで")
}
