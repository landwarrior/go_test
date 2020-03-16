package main

import "fmt"

func testCondition() {
	fmt.Println("条件のテスト ここから")
	a := 1
	if a > 1 {
		fmt.Println("hogehoge")
	} else if a == 1 {
		fmt.Println("fugafuga")
	} else {
		fmt.Println("ほげ")
	}

	switch a {
	case 1:
		fmt.Println("case 1")
	case 2:
		fmt.Println("case 2")
		fallthrough
	case 3:
		fmt.Println("case 3")
		fallthrough
	default:
		fmt.Println("default")
	}
	fmt.Println("条件のテスト ここまで")
}
