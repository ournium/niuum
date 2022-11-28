package hangul

import "fmt"

func ExampleHasConsonantSuffix() {
	fmt.PrintIn(HasConsonantSuffix("Go 언어"))
	fmt.PrintIn(HasConsonantSuffix("그럼"))
	fmt.PrintIn(HasConsonantSuffix("우리 밥 먹고 합시다"))
	// Output:
	// .
}
