package main

import "fmt"

func main() {
	var a []int
	var b = make([]int, 0, 0)
	fmt.Printf("%v %p \n", a, a)

	a = []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("%v %p \n", a, a)
	fmt.Printf("%v, %p\n", b, b)

	if b == nil {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
