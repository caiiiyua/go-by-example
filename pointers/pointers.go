package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	n := 5
	fmt.Println(n)
	zeroval(n)
	fmt.Println("zeroval: ", n)
	zeroptr(&n)
	fmt.Println("zeroptr: ", n)
	fmt.Println("address: ", &n)
}
