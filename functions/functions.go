package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func main() {
	r := plus(1, 2)
	fmt.Println("1 + 2 =", r)
}
