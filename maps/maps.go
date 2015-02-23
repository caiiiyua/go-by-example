package main

import "fmt"

func main() {
	m := make(map[string]int)
	fmt.Println("empty:", m)

	m["k1"] = 7
	m["k2"] = 11
	fmt.Println("set:", m)
	fmt.Println("get:", m["k2"])

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("del:", m)
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
