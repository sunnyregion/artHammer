package golint

import (
	"fmt"
)

func golint_test() {
	fmt.Println("test string!")

	var str string = "abc" //golint check not right

	fmt.Println(str)

	var x int
	x += 1 //golint check not right
}
