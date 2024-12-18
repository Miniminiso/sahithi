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
	array1 := make([]string, 8)
	array1[0] = "sahi"
	array1[1] = "swetha"
	array1[3] = "harini"
	fmt.Println(array1)
}
