package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(a string) (int, string) {
	return len(a), strings.ToUpper(a)
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
