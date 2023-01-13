package main

import "fmt"

func main() {
	slice1 := make([]int, 3, 5)
	// slice1 = [0 0 0]
	slice2 := append(slice1, 4, 5)
	// slice2 = [0 0 0 4 5]
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))
	fmt.Printf("slice1 %p, slice2 %p \n", slice1, slice2)

	slice1[1] = 100
	// slice = [0 100 0]
	// slice = [0 100 0 4 5]
	fmt.Println("After change second element")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))
	fmt.Printf("slice1 %p, slice2 %p \n", slice1, slice2)

	slice1 = append(slice1, 500)
	// slice1 = [0 100 0 500]
	// slice2 = [0 100 0 500 5]
	fmt.Println("After append 500")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))
	fmt.Printf("slice1 %p, slice2 %p \n", slice1, slice2)
}
