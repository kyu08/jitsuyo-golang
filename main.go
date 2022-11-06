package main

import (
	"fmt"
)

const (
	a = iota + 1 // 1
	b            // 2
	c            // 3
	_            // 4だが使われない
	e            // 5
	f = iota     // 6 iotaの値はリセットされない
)

const (
	g = iota // 再び0になる
	h        // 1
	i        // 2
)

func main() {
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
	fmt.Printf("e: %v\n", e)
	fmt.Printf("f: %v\n", f)
}
