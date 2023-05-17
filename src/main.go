package main

import (
	"fmt"
)

type TestStruct struct {
	hoge string
	huga string
}

func (test *TestStruct) modify(i string, j string) {
	test.hoge = i
	test.huga = j
}

func main() {
	test_1 := TestStruct{
		hoge: "a",
		huga: "b",
	}
	fmt.Println(test_1)
	test_1.modify("c", "d")
	fmt.Println(test_1)
}
