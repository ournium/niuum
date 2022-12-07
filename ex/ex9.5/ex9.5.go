package main

import "fmt"

func HasRichFriend() bool {
	return true
}

func GetFriendsCount() int {
	return 3
}

func main() {
	price := 35000

	if price > 50000 {
		if HasRichFriend() {
		} else {
			fmt.Println("나눠내자")
		}
	} else if price >= 30000 && price <= 50000 {
		if GetFriendsCount() {
			fmt.Println("아이쿠 신발끈이..")
		} else {
			fmt.Println("사람도 얼마 없는데 나눠 내자")
		}
	} else {
		fmt.Println("오늘은 내가 쏜다")
	}
}
