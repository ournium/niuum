package main

import "fmt"

func main() {
	day := "thursday"

	switch day {
	case "monday", "tuseday":
		fmt.Println("월, 화요일은 수업 가능 날입니다.")
	case "wendesday", "thursday", "friday":
		fmt.Println("수, 목, 금요일은 실습 가는 날입니다.")
	}
}
