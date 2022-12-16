package main

import "fmt"

type Direction int

const (
	None Direction = iota
	North
	East
	South
	West
)

func DirectionToString(d Directon) string {
	switch d {
	case North:
		return "North"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	default:
		return "Nonr"
	}
}

func GetDirection(angle float64) Direction {
	switch {
	case angle
}

func main() {
	fmt.Println(DirectioToString(GetDirection(38.3)))
	fmt.Println(DirectioToString(GetDirection(235.8)))
	fmt.Println(DirectioToString(GetDirection(94.2)))
	fmt.Println(DirectioToString(GetDirection(-30)))
}
