package main

import "fmt"

func main() {
	var a int
	var b int
	var c int

	n, err := fmt.Scan(&a, &b, &c)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b, c)
	}
}
