package main

import "fmt"

func main() {
	light := "red"

	if light == "green" {
		fmt.Println("길을건넌다")
	} else {
		fmt.Println("대기한다")
	}
}
