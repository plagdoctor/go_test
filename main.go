package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(a string) (wordLength int, wordUpper string) {
	defer fmt.Println("done.")
	wordLength = len(a)
	wordUpper = strings.ToUpper(a)

	return

}
func repeat(words ...string) {
	fmt.Println(words)
}

func superAdd(numbers ...int) int {
	total := 0
	for i := 1; i <= 6; i++ {
		println(i)
	}

	for _, number := range numbers {
		total += number
	}
	println(total)
	return 1
}
func canIDrink(age int) bool {
	switch koreanAge := age; koreanAge {
	case 18:
		return true
	case 10:
		return true
	}
	return false

}
func main() {
	/*
		fmt.Println("dd")

		var name string = "jinyoung"
		name = "nico"

		fmt.Println(name)

		name_t_lengh, name_t_upper := lenAndUpper(name)
		fmt.Println(name_t_lengh, name_t_upper)
		repeat("mike", "jinyoung", "haha", "ilence")
	*/
	//superAdd(1, 2, 3, 4, 5, 6)

	println(canIDrink(18))
	println(canIDrink(21))
}
