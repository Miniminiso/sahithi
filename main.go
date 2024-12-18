package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	test2()
}

func test2() {
	array1 := make([]int, 8)
	array1[0] = 2
	array1[1] = 9
	array1[3] = 23
	fmt.Println(array1)
}
