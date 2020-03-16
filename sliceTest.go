package main

import "fmt"

func testSlice() {
	fmt.Println("スライスのテスト ここから")
	var a []int
	b := []string{"hoge", "fuga"}
	c := []int{1, 2}
	fmt.Printf("type:%T value:%v len:%d cap:%d\n", a, a, len(a), cap(a))
	fmt.Printf("type:%T value:%v len:%d cap:%d\n", b, b, len(b), cap(b))
	fmt.Printf("type:%T value:%v len:%d cap:%d\n", c, c, len(c), cap(c))

	a = append(c, 1)
	fmt.Printf("type:%T value:%v len:%d cap:%d\n", a, a, len(a), cap(a))
	a = append(a, 3)
	fmt.Printf("type:%T value:%v len:%d cap:%d\n", a, a, len(a), cap(a))
	a = append(a, 4)
	fmt.Printf("type:%T value:%v len:%d cap:%d\n", a, a, len(a), cap(a))
	a = append(a, 5, 6, 7, 8)
	fmt.Printf("type:%T value:%v len:%d cap:%d\n", a, a, len(a), cap(a))
	fmt.Println(a[3:])
	fmt.Println("スライスのテスト ここまで")
}
