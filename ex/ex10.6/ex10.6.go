package main

import "fmt"

func getMyage() int {
	return 22
}

func main() {
	switch age := getMyage(); age {
	case 10:
		fmt.Println("Teenage")
	case 33:
		fmt
