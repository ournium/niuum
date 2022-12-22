package main

import "fmt"

func PrintAvg(name string, math int, eng int, history int) {
	total := math + eng + history
	avg := total / 3
	fmt.Println(name, "님 평균 점수는", avg, "입니다.")
}

func main() {
	PrintAvg("김일등", 80, 74, 95)
	PrintAvg("송이등", 88, 92, 53)
	PrintAvg("박삼등", 78, 73, 78)
}
