package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func main() {

	var sumOfTwo = sum(1, 2)
	fmt.Println(sumOfTwo)
}
