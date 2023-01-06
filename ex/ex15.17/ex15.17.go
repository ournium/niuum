package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var str string = "Hello"
	stringheader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	addr1 := stringheader.Data

	str += " world"
	addr2 := stringheader.Data
	fmt.Println(str)

	str += " Welcome!"
	addr3 := stringheader.Data

	fmt.Println(str)
	fmt.Printf("addr1:\t%x\n", addr1)
	fmt.Printf("addr2:\t%x\n", addr2)
	fmt.Printf("addr3:\t%x\n", addr3)
}
