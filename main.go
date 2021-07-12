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
func main() {
	fmt.Println("dd")

	var name string = "jinyoung"
	name = "nico"

	fmt.Println(name)

	name_t_lengh, name_t_upper := lenAndUpper(name)
	fmt.Println(name_t_lengh, name_t_upper)
	repeat("mike", "jinyoung", "haha", "ilence")

}
