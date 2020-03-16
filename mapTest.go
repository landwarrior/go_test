package main

import "fmt"

func testMapping() {
	fmt.Println("マップのテスト ここから")
	var test1 map[string]int
	fmt.Printf("type:%T value:%v\n", test1, test1)
	test2 := map[string][]int{"hoge": []int{1, 2, 3}}
	fmt.Printf("type:%T value:%v\n", test2, test2)
	fmt.Println(test2["hoge"][1])
	test3 := map[string][3]int{"fuga": [3]int{1, 2, 3}}
	fmt.Printf("type:%T value:%v\n", test3, test3)
	fmt.Println(test3["fuga"][1])
	fmt.Println(test2["hoge"][1] == test3["fuga"][1])
	fmt.Println("マップのテスト ここまで")
}
